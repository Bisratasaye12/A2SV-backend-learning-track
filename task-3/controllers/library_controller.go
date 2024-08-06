package controllers

import (
	"Task-3/models"
	"Task-3/services"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	fmt.Println("7. Add Member")
	fmt.Println("8. Remove Member")
	fmt.Println("9. Exit")
}

func bookInput() models.Book {
	newBook := models.Book{Status: "Available"}
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
	// fmt.Println("The library books",library.Books)

}

func RemoveBook() {
	var memberID int
	var err error
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Enter book ID to remove:")
		scanner.Scan()
		input := scanner.Text()
		memberID, err = strconv.Atoi(input)

		if err != nil {
			fmt.Println("Invalid book ID")
			continue
		}
		break
	}
	// Call the RemoveBook method from the library service
	library.RemoveBook(memberID)
	// fmt.Println("The library books",library.Books)
}

// member related functions
func AddMember() {
	newMember := models.Member{}
	newMember.BorrowedBooks = []models.Book{}
	newMember.ID = library.NextMemberID
	library.NextMemberID++

	// Prompt the user to enter member name
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter member name:")
	scanner.Scan()
	newMember.Name = scanner.Text()
	newMember.Name = strings.TrimSpace(newMember.Name)
	library.AddMember(newMember)
	fmt.Println("The library members", library.Members)
}

func RemoveMember() {
	var memberID int
	var err error
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Enter book ID to remove:")
		scanner.Scan()
		input := scanner.Text()
		memberID, err = strconv.Atoi(input)

		if err != nil {
			fmt.Println("Member ID is invalid")
			continue
		}
		break
	}
	// Call the RemoveBook method from the library service
	library.RemoveMember(memberID)
}

func BorrowBook() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter book ID to borrow:")
	
	for _, book := range library.Books {
		fmt.Println("ID:", book.ID, "Title:", book.Title, "Author:", book.Author)
	}

	scanner.Scan()
	bookID, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Invalid book ID")
		return
	}

	fmt.Println("Enter member ID:")
	scanner.Scan()
	memberID, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Invalid member ID")
		return
	}
	if _, ok := library.Members[memberID]; !ok {
		fmt.Println("Member not found")
		return
	}

	if _, ok := library.Books[bookID]; !ok {
		fmt.Println("book not found")
		return
	}

	library.BorrowBook(bookID, memberID)
}

func ReturnBook() {
	scanner := bufio.NewScanner(os.Stdin)
	var bookID, memberID int
	var err error

	for {
		fmt.Println("Enter book ID to return:")
		scanner.Scan()
		bookID, err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Invalid book ID")
			continue
		}

		fmt.Println("Enter member ID:")
		scanner.Scan()
		memberID, err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Invalid member ID")
			continue
		}

		break
	}

	if _, ok := library.Books[bookID]; !ok {
		fmt.Println("Book not found")
		return
	}
	if _, ok := library.Members[memberID]; !ok {
		fmt.Println("Member not found")
		return
	}

	library.ReturnBook(bookID, memberID)
}

func ListAvailableBooks() {
	books := library.ListAvailableBooks()
	fmt.Println("------- Available books -------")
	for _, book := range books {
		fmt.Println("ID:", book.ID, "Title:", book.Title, "Author:", book.Author)
	}
}

func ListBorrowedBooks() {
	scanner := bufio.NewScanner(os.Stdin)
	var memberID int
	var err error
	for {
		fmt.Println("Enter member ID:")
		scanner.Scan()
		memberID, err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Invalid member ID")
			continue
		}
		break
	}

	books := library.ListBorrowedBooks(memberID)
	if len(books) == 0 {
		fmt.Println("No borrowed books")
		return
	}
	fmt.Println("------- Borrowed books -------")
	for _, book := range books {
		fmt.Println(book)
	}
}

// func MyTest(){
// 	fmt.Println("library books",library.Books, "\n")
// 	fmt.Println("library members",library.Members, "\n")
	
// 	for _, member := range library.Members{
// 		fmt.Println("member borrowed books",member.BorrowedBooks, "\n")
// 	}
// }
