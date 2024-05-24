package main

import (
	"log"

	"gorm.io/driver/sqlite"

	"gorm.io/gorm"

)

var DB *gorm.DB

// InitializeDB initializes the SQLite database and performs auto-migration for the Book struct.
func InitializeDB() {

	var err error

	DB, err = gorm.Open(sqlite.Open("books.sqlite"), &gorm.Config{})
	
	if err != nil {
	
		log.Fatal("Error opening database connection: ", err)
	
	}

	// Auto-migrate the schema

	err = DB.AutoMigrate(&Book{})

	if err != nil {

		log.Fatal("Error auto-migrating database schema: ", err)

	}

	log.Println("Database initialized")
}

// CloseDB closes the database connection.
func CloseDB() {

	if DB != nil {

		db, err := DB.DB()


		if err != nil {

			log.Println("Error getting database connection: ", err)

			return
		}

		db.Close()

		log.Println("Database connection closed")

	}

}
