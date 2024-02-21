package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "2004"
	dbname   = "library"
)

type Book struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Price     uint
	Available bool
}

func main() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Book{})

	db.Create(&Book{Name: "Abay zholy", Price: 5000, Available: true})

	var book Book

	db.First(&book, 1)

	fmt.Print(book)
}
