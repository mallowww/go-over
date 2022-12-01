package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

const cooursePath = "courses"
const basePath = "/api"

type Course struct {
	CourseID   int     `json: "courseid"`
	Coursename string  `json: "coursename"`
	Price      float64 `json: "price"`
	ImageURL   string  `json: "imageurl"`
}

func SetupDB() {
	var err error
	DB, err = sql.Open("mysql", "root:mysqlbro124@tcp(127.0.0.1:3306)/coursedb")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(DB)
	DB.SetConnMaxLifetime(time.Minute * 2)
	DB.SetMaxIdleConns(10)
	DB.SetConnMaxIdleTime(10)
}

func getCourseList() ([]Course, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	res, err := DB.QueryContext(ctx, `SELECT 
	courseid,
	coursename,
	price,
	image_url
	FROM courseonline`)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer res.Close()
	courses := make([]Course, 0)
	for res.Next() {
		var course Course
		res.Scan(&course.CourseID,
			&course.Coursename,
			&course.Price,
			&course.ImageURL,
		)

		courses = append(courses, course)
	}
	return courses, nil
}

func handleCourses(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		courseList, err := getCourseList()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		j, err := json.Marshal(courseList)
		if err != nil {
			log.Fatal(err)
		}
		_, err = w.Write(j)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// Enable CORS in Golang
// https://stackoverflow.com/questions/39507065/enable-cors-in-golang
// https://stackoverflow.com/questions/69608544/enable-cors-policy-in-golang
func corsMiddleware(handler http.Handler) http.Handler{
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-rigin", "*")
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Add("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token")
		handler.ServeHTTP(w, r)
	})
}

func SetupRoutes(apiBasePath string)  {
	courseHandler := http.HandlerFunc(handleCourses)
	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, cooursePath), corsMiddleware(courseHandler))
	
}

func main() {
	SetupDB()
	SetupRoutes(basePath)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
