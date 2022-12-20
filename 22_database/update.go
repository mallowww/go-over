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

	statement, err := db.Prepare("UPDATE users SET name=$2, age=$3 WHERE id=$1;")
	if err != nil {
		log.Fatal("can't prepare update statement ", err)
	}

	_, err = statement.Exec(1, "sveltekid", 17)
	if err != nil {
		log.Fatal("can't excute update ", err)
	}

	fmt.Println("update success")
}
