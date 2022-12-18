package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
<<<<<<< HEAD
<<<<<<< HEAD
	"time"
=======
>>>>>>> 55103d7 (:construction: feat(middleware): fcf, fl, hof, logMdw)
=======
	"time"
>>>>>>> a84c37a (:sparkles: feat(middleware): log)
)

type User struct {
	ID   int
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = []User{
	{ID: 1, Name: "okiebro", Age: 15},
	{ID: 2, Name: "nuxt", Age: 16},
	{ID: 3, Name: "sveltekit", Age: 17},
}

func usersHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		b, err := json.Marshal(users)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.Write(b)
		return
	}

	if req.Method == "POST" {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		var u User
		err = json.Unmarshal(body, &u)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		users = append(users, u)

		// w.Write([]byte(`{"name": "okiela", "method":"POST"}`))
		fmt.Fprintf(w, "ok this create users - %s", "POST")
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> a84c37a (:sparkles: feat(middleware): log)
// func logMiddleware(Handler http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		start := time.Now()
// 		Handler.ServeHTTP(w, r)
// 		log.Printf("Server http middleware: %s %s %s %s ", r.RemoteAddr, r.Method, r.URL, time.Since(start))
// 	}
// }

type Logger struct {
	Handler http.Handler
}

func (l Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	l.Handler.ServeHTTP(w, r)
	log.Printf("Server http middleware: %s %s %s %s ", r.RemoteAddr, r.Method, r.URL, time.Since(start))

}

<<<<<<< HEAD
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/users", usersHandler)
	mux.HandleFunc("/health", healthHandler)

	logMux := Logger{Handler: mux}
	srv := http.Server{
		Addr:    ":8080",
		Handler: logMux,
	}

	log.Println("server started at default port, :8080")
	log.Fatal(srv.ListenAndServe())
=======
=======
>>>>>>> a84c37a (:sparkles: feat(middleware): log)
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/users", usersHandler)
	mux.HandleFunc("/health", healthHandler)

	logMux := Logger{Handler: mux}
	srv := http.Server{
		Addr:    ":8080",
		Handler: logMux,
	}

	log.Println("server started at default port, :8080")
<<<<<<< HEAD
	log.Fatal(http.ListenAndServe(":8080", nil))
>>>>>>> 55103d7 (:construction: feat(middleware): fcf, fl, hof, logMdw)
=======
	log.Fatal(srv.ListenAndServe())
>>>>>>> a84c37a (:sparkles: feat(middleware): log)
	log.Println("close")
}
