
# Library Management System

## Overview

The Library Management System (LMS) is a console-based application implemented in Go. It allows users to manage books and members in a library, supporting operations such as adding and removing books and members, borrowing and returning books, and listing available and borrowed books.

## File Structure

```
library_management/
├── main.go
├── controllers/
│   └── library_controller.go
├── models/
│   └── book.go
│   └── member.go
├── services/
│   └── library_service.go
├── docs/
│   └── documentation.md
└── go.mod
```

## Files Description

### `main.go`

The entry point of the application. It displays a welcome message, presents a menu to the user, and handles user input to perform various actions based on the user's choice.

### `services/library_service.go`

Defines the `LibraryManager` interface and its implementation, `Library`. It manages books and members, handling operations such as adding, removing, borrowing, and returning books.

### `models/book.go`

Defines the `Book` struct with fields for ID, title, author, and status.

### `models/member.go`

Defines the `Member` struct with fields for ID, name, and a list of borrowed books.

### `controllers/library_controller.go`

Contains functions for interacting with the user. It handles user input, calls service methods, and displays results.

## Setup Instructions

1. Clone the repository.
2. Navigate to the project directory.
3. Run `go mod tidy` to ensure all dependencies are downloaded.
4. Execute `go run main.go` to start the application.

## Usage

- **Add Book:** Enter details to add a new book.
- **Remove Book:** Provide the book ID to remove it.
- **Borrow Book:** Specify the book and member IDs to borrow a book.
- **Return Book:** Specify the book and member IDs to return a book.
- **List Available Books:** View all books that are currently available.
- **List Borrowed Books:** View all books borrowed by a specific member.
- **Add Member:** Add a new member to the library.
- **Remove Member:** Remove a member using their ID.

## Contributing


## License
```
