package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	url := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal("can't connect to database", err)
	}
	defer db.Close()

	row := db.QueryRow("INSERT INTO users (name,age) values ($1, $2) RETURNING id", "okie", 15)
	var id int
	err = row.Scan(&id)
	if err != nil {
		log.Fatal("can't insert data", err)
	}
	fmt.Println("insert todo success id: ", id)
}
