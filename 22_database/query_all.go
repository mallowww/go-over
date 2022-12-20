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
		log.Fatal("can't connect to db", err)
	}
	defer db.Close()

	statement, err := db.Prepare("SELECT id, name, age FROM users")
	if err != nil {
		log.Fatal("can't prepare query all users", err)
	}

	rows, err := statement.Query()
	if err != nil {
		log.Fatal("can't query all users", err)
	}

	for rows.Next() {
		var (
			id   int
			name string
			age  int
		)
		err := rows.Scan(&id, &name, &age)
		if err != nil {
			log.Fatal("can't scan row into variable", err)
		}
		fmt.Println(id, name, age)
	}
	fmt.Println("query all users success")
}
