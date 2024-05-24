package main

import (

	"log"

	"net/http"

	"strconv"

	"github.com/labstack/echo/v4"
)

// Book represents a book in the database.




type Book struct {

	ID          int    `json:"id"`

	Title       string `json:"title"`

	Author      string `json:"author"`

	Description string `json:"description"`

	Precis      int    `json:"precis"`
}

// Main function to start the server.

func main() {

	InitializeDB()  // Initialize the database

	defer CloseDB() // Close the database connection when main function exits

	e := echo.New()

	// Routes for API
	e.GET("/api/books", listBooks)

	e.POST("/api/books", addBook)

	e.DELETE("/api/books/:id", deleteBook)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// Handler functions for CRUD operations

func listBooks(c echo.Context) error {

	var books []Book

	DB.Find(&books)

	return c.JSON(http.StatusOK, books)
}

func addBook(c echo.Context) error {

	var book Book

	if err := c.Bind(&book); err != nil {

		log.Println("Error binding request body:", err)

		return err

	}

	log.Printf("Received book data: %+v\n", book)

	DB.Create(&book)

	return c.JSON(http.StatusCreated, book)
}

func deleteBook(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {

		return c.String(http.StatusBadRequest, "Invalid ID")

	}

	var book Book

	DB.First(&book, id)

	if book.ID == 0 {

		return c.String(http.StatusNotFound, "Book not found")
	}

	DB.Delete(&book)

	return c.NoContent(http.StatusNoContent)
}
