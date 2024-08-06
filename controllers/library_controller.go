package controllers

import (
	"Task-3/models"
	"bufio"
	"fmt"
	"os"
	"Task-3/services"
)

var library = services.NewLibrary()

// Utility functions
func ShowMenu() {
	fmt.Println("1. Add Book")
	fmt.Println("2. Remove Book")
	fmt.Println("3. Borrow Book")
	fmt.Println("4. Return Book")
	fmt.Println("5. List Available Books")
	fmt.Println("6. List Borrowed Books")
	fmt.Println("7. Exit")
}

func bookInput() models.Book {
	newBook := models.Book{}
	newBook.ID = library.NextBookID
	// Prompt the user to enter title and Author
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter book title:")
	scanner.Scan()
	newBook.Title = scanner.Text()
	fmt.Println("Enter book author:")
	scanner.Scan()
	newBook.Author = scanner.Text()

	return newBook
}



// Main logic functions

func AddBook() {
	newBook := bookInput()

	// Call the AddBook method from the library service
	library.AddBook(newBook)
	fmt.Println("The library books",library.Books)

}