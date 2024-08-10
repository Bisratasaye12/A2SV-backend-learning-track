package services

import (
	"Task-3/models"
	"testing"
)

func TestLibrary(t *testing.T) {
	lib := NewLibrary()

	// Test AddBook
	book := models.Book{ID: 1, Title: "Test Book", Author: "Test Author", Status: "Available"}
	lib.AddBook(book)
	if len(lib.Books) != 1 {
		t.Errorf("Expected 1 book, got %d", len(lib.Books))
	}
	if lib.Books[1] != book {
		t.Errorf("Expected book %v, got %v", book, lib.Books[1])
	}

	// Test RemoveBook
	lib.RemoveBook(1)
	if len(lib.Books) != 0 {
		t.Errorf("Expected 0 books, got %d", len(lib.Books))
	}

	// Test BorrowBook
	lib.AddBook(book)
	member := models.Member{ID: 1, Name: "Test Member", BorrowedBooks: []models.Book{}}
	lib.AddMember(member)
	err := lib.BorrowBook(1, 1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if len(lib.Members[1].BorrowedBooks) != 1 {
		t.Errorf("Expected 1 borrowed book, got %d", len(lib.Members[1].BorrowedBooks))
	}

	// Test ReturnBook
	err = lib.ReturnBook(1, 1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if len(lib.Members[1].BorrowedBooks) != 0 {
		t.Errorf("Expected 0 borrowed books, got %d", len(lib.Members[1].BorrowedBooks))
	}

	// Test AddMember
	newMember := models.Member{ID: 2, Name: "New Member", BorrowedBooks: []models.Book{}}
	lib.AddMember(newMember)
	if len(lib.Members) != 2 {
		t.Errorf("Expected 2 members, got %d", len(lib.Members))
	}

	// Test RemoveMember
	lib.RemoveMember(2)
	if len(lib.Members) != 1 {
		t.Errorf("Expected 1 member, got %d", len(lib.Members))
	}
}
