package services

import (
	"Task-3/models"
	"fmt"
)

type LibraryManager interface{
	AddBook(book models.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
}

type Library struct {
	Books map[int]models.Book
	Members map[int]models.Member
	NextBookID int
}

func NewLibrary() *Library {
    return &Library{
        Books:      make(map[int]models.Book),
        Members:    make(map[int]models.Member),
        NextBookID: 1, 
    }
}
func (lib *Library) AddBook(book models.Book) {
	lib.Books[book.ID] = book
	lib.NextBookID++
	fmt.Println("Book added successfully")
}