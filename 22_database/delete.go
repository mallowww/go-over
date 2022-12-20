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
		log.Fatal("can't connect to db ", err)
	}
	defer db.Close()

	statement, err := db.Prepare("DELETE FROM users WHERE id=$1")
	if err != nil {
		log.Fatal("can't prepare delete statment ", err)
	}

	if _, err := statement.Exec(1); err != nil {
		log.Fatal("can't execute delete statement ", err)
	}
	fmt.Println("delete success")
}
