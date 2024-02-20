package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "2004"
	dbname   = "library"
)

func main() {
	// Строка подключения
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Открываем подключение к базе данных
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	for {
		fmt.Println("Task Manager Menu:")
		fmt.Println("1. Add Task")
		fmt.Println("2. View Tasks")
		fmt.Println("3. Mark task as Complete")
		fmt.Println("4. Remove Task")
		fmt.Println("5. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		if _, err := fmt.Scan(&choice); err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
	}

}
