package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/lib/pq"
)

type User struct {
	ID   int
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK")
}

func getAllUser(c echo.Context) error {
	statement, err := db.Prepare("SELECT id, name, age FROM users")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't prepare query all users statement"})
	}
	rows, err := statement.Query()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't query all users:" + err.Error()})
	}

	users := []User{}
	for rows.Next() {
		var u User
		err = rows.Scan(&u.ID, &u.Name, &u.Age)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, Err{Message: "can't scan user:" + err.Error()})
		}
		users = append(users, u)
	}

	return c.JSON(http.StatusOK, users)
}

type Err struct {
	Message string `json:"message"`
}

func createUser(c echo.Context) error {
	var u User
	err := c.Bind(&u)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}

	rows := db.QueryRow("INSERT INTO users (name, age) values ($1,$2) RETURNING id", u.Name, u.Age)
	err = rows.Scan(&u.ID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, u)
}

var db *sql.DB

func main() {
	var err error
	url := os.Getenv("DATABASE_URL")
	db, err = sql.Open("postgres", url)
	if err != nil {
		log.Fatal("can't connect to database ", err)
	}
	defer db.Close()

	createTable := `CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name TEXT, age INT);`
	if _, err = db.Exec(createTable); err != nil {
		log.Fatal("can't execute 'create table' ", err)
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/health", healthHandler)

	e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "okiebro" && password == "welp4455" {
			return true, nil
		}
		return false, nil
	}))

	e.GET("/users", getAllUser)
	e.POST("/users", createUser)

	log.Println("server started at default port, :8080")
	log.Fatal(e.Start(":8080"))
	log.Println("close")
}
