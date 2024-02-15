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

	// Проверяем подключение
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Успешное подключение к базе данных")
}
