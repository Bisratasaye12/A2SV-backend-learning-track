package main

import (
	"fmt"
	"Task-3/controllers"
)


func main(){
	fmt.Println("Welcome to the Library Management System")
	controllers.ShowMenu()

	for {
		var choice int
		fmt.Println("Enter your choice:")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			controllers.AddBook()
		case 2:
			controllers.RemoveBook()
		// case 3:
		// 	controllers.BorrowBook()
		// case 4:
		// 	controllers.ReturnBook()
		// case 5:
		// 	controllers.ListAvailableBooks()
		// case 6:
		// 	controllers.ListBorrowedBooks()
		case 7:
			controllers.AddMember()
		case 8:
			controllers.RemoveMember()
		case 9:
			fmt.Println("Exiting the program")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}