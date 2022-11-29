package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func createTable(db *sql.DB) {
	query := `CREATE TABLE users (
		id INT AUTO_INCREMENT,
		username TEXT NOT NULL,
		password TEXT NOT NULL,
		created_at DATETIME,
		PRIMARY KEY (id)
		);`

	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}
}

func Insert(db *sql.DB) {
	var (
		username string
		password string
	)
	fmt.Scan(&username)
	fmt.Scan(&password)
	createdAt := time.Now()

	res, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?,?,?)`, username, password, createdAt)
	if err != nil {
		log.Fatal(err)
	}
	id, err := res.LastInsertId()
	fmt.Println(id)
}

func query(db *sql.DB) {
	var (
		id         int
		coursename string
		price      float64
		instructor string
	)

	fmt.Println("- Input ID below -")
	var searchID int
	fmt.Scan(&searchID)
	query := "SELECT id, coursename, price, instructor FROM onlinecourse WHERE id = ?"
	if err := db.QueryRow(query, searchID).Scan(&id, &coursename, &price, &instructor); err != nil {
		log.Fatal(err)
	}
	fmt.Println(id, coursename, price, instructor)
}

func delete(db *sql.DB) {
	var deleteID int
	fmt.Scan(&deleteID)
	_, err := db.Exec(`DELETE FROM users WHERE id = ?`, deleteID)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	db, err := sql.Open("mysql", "root:mysqlbro124@tcp(127.0.0.1:3306)/coursedb")
	if err != nil {
		fmt.Println("failed to connect")
	} else {
		fmt.Println("connect successfully")
	}
	// createTable(db)
	// Insert(db)
	delete(db)
}
