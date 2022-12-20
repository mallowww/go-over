package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func init() {
	fmt.Println("main init")
}

func main() {
	url := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal("connect to database error", err)
	}

	defer db.Close()

	createTable := `CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name TEXT, age INT);`
	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatal("can't create table", err)
	}

	fmt.Println("create table success")
}
