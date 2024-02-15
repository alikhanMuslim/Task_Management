package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/alikhanMuslim/Task_Management/pkg/dbs"
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

	fmt.Println(dbs.GetTaskByID(db, 1))

	fmt.Println("Успешное подключение к базе данных")
}
