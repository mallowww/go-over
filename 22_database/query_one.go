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

	statement, err := db.Prepare("SELECT id, name, age FROM users where id=$1")
	if err != nil {
		log.Fatal("can't prepare query one row")
	}

	rowId := 2
	row := statement.QueryRow(rowId)
	var id, age int
	var name string

	err = row.Scan(&id, &name, &age)
	if err != nil {
		log.Fatal("can't scan row into var", err)
	}

	fmt.Println("one row of: ", id, name, age)
}
