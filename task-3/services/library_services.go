package services

import (
	"Task-3/models"
	"fmt"
)

type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
	AddMember(member models.Member)
	RemoveMember(id int)
}


type Library struct {
	Books      map[int]models.Book
	Members    map[int]models.Member
	NextBookID int
	NextMemberID int
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

func (lib *Library) RemoveBook(bookID int) {
	if _, ok := lib.Books[bookID]; !ok {
		fmt.Println("Book not found")
		return
	}

	if lib.Books[bookID].Status == "Borrowed" {
		fmt.Println("Book is borrowed, cannot remove")
		return
	}
	
	delete(lib.Books, bookID)
	fmt.Println("Book removed successfully")
}

// func (lib *Library) BorrowBook(bookID int, memberID int) error {
// 	book, ok := lib.Books[bookID]
// 	if !ok {
// 		return fmt.Errorf("Book not found")
// 	}
// 	if book.Status == "Borrowed" {
// 		return fmt.Errorf("Book is already borrowed")
// 	}
// 	book.Status = "Borrowed"
// 	member := lib.Members[memberID]
// 	return nil
// }

func (lib *Library) AddMember(newMember models.Member) {
	lib.Members[newMember.ID] = newMember
	fmt.Println("Member added successfully\n")
}

func (lib *Library) RemoveMember(id int) {
	if _, ok := lib.Members[id]; !ok {
		fmt.Println("Member not found")
		return
	}
	delete(lib.Members, id)
	fmt.Println("Member removed successfully\n")
}