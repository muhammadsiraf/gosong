package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// var baseURL = "http://localhost:8000"

type pelajar1 struct {
	id    string
	name  string
	age   int
	grade int
}

const (
	host     = "localhost"
	port     = "5432"
	userdb   = "postgres"
	password = "Sindoro10"
	dbname   = "student"
)

func connect() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", host, port, userdb, password, dbname)
	// connStr := "postgres:Sindoro10@localhost/student?sslmode=disable"
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func sqlQuery() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("select name, grade from student where id = $1")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var result1 = pelajar1{}
	stmt.QueryRow("E001").Scan(&result1.name, &result1.grade)
	fmt.Printf("name: %s\ngrade: %d\n", result1.name, result1.grade)

	var result2 = pelajar1{}
	stmt.QueryRow("W001").Scan(&result2.name, &result2.grade)
	fmt.Printf("name: %s\ngrade: %d\n", result2.name, result2.grade)

	var result3 = pelajar1{}
	stmt.QueryRow("B001").Scan(&result3.name, &result3.grade)
	fmt.Printf("name: %s\ngrade: %d\n", result3.name, result3.grade)
}

func main() {
	sqlQuery()
}
