package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

func main() {
	http.HandleFunc("/users", func(w http.ResponseWriter, req *http.Request) {
		if req.Method == "GET" {
			// w.Write([]byte(`{"name": "okiela", "method":"GET"}`))
			log.Println("GET")
			b, err := json.Marshal(users)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				// w.Write([]byte(`"500 - Something bad happended!`))
				w.Write([]byte(err.Error()))
				return
			}
			w.Write(b)
			return
		}

		if req.Method == "POST" {
			log.Println("POST")

			body, err := ioutil.ReadAll(req.Body)
			if err != nil {
				// fmt.Fprintf(w, "error: %v", err)
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
	})

	log.Println("server started at default port, :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
	log.Println("close")
}
