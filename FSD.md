# Functional Specification Document (FSD)

## Application: Library Management System (LMS)

### 1. Introduction

The **Library Management System (LMS)** is a Go-based application designed to manage books in a library. It allows users to perform CRUD (Create, Read, Update, Delete) operations on books, including adding new books, viewing book details, updating book information, and deleting books. The application will be a command-line tool that interacts with an in-memory database.

---

### 2. Objectives

- Provide a simple interface to manage books in a library.
- Allow users to perform CRUD operations on books.
- Store book information in an in-memory data structure.
- Ensure the application is easy to use and maintain.

---

### 3. Scope

The application will be a command-line tool with the following features:

- Add a new book.
- View all books.
- View details of a specific book.
- Update an existing book.
- Delete a book.

---

### 4. Functional Requirements

#### 4.1. Add a New Book

- **Description**: The user should be able to add a new book by providing the book title, author, publication year, and ISBN.
- **Input**: Book title (string), author (string), publication year (int), ISBN (string).
- **Output**: Confirmation message that the book has been added.

#### 4.2. View All Books

- **Description**: The user should be able to view a list of all books.
- **Input**: None.
- **Output**: List of all books with their titles and IDs.

#### 4.3. View Book Details

- **Description**: The user should be able to view the details of a specific book by providing the book ID.
- **Input**: Book ID (int).
- **Output**: Book details including title, author, publication year, and ISBN.

#### 4.4. Update a Book

- **Description**: The user should be able to update the details of an existing book by providing the book ID and new details.
- **Input**: Book ID (int), new title (string), new author (string), new publication year (int), new ISBN (string).
- **Output**: Confirmation message that the book has been updated.

#### 4.5. Delete a Book

- **Description**: The user should be able to delete a book by providing the book ID.
- **Input**: Book ID (int).
- **Output**: Confirmation message that the book has been deleted.

---

### 5. Non-Functional Requirements

- **Performance**: The application should respond to user commands within 1 second.
- **Usability**: The command-line interface should be intuitive and easy to use.
- **Maintainability**: The code should be well-structured and documented for easy maintenance.

---

### 6. Data Model

The application will use a simple in-memory data structure to store book information. Each book will have the following fields:

- **ID**: Unique identifier for the book (int).
- **Title**: Title of the book (string).
- **Author**: Author of the book (string).
- **PublicationYear**: Year the book was published (int).
- **ISBN**: International Standard Book Number (string).

---

### 7. User Interface

The application will be a command-line tool with the following commands:

- `add <title> <author> <publicationYear> <ISBN>`: Add a new book.
- `list`: List all books.
- `view <id>`: View details of a specific book.
- `update <id> <title> <author> <publicationYear> <ISBN>`: Update a book.
- `delete <id>`: Delete a book.

---

### 8. Error Handling

- The application should handle invalid inputs gracefully and provide meaningful error messages.
- If a book ID is not found, the application should inform the user.

---

### 9. Assumptions

- The application will not persist data; all data will be lost when the application is closed.
- The application will be used by a single user at a time.

---

### 10. Dependencies

- Go programming language (version 1.20 or higher).
- Standard Go libraries (no external dependencies).

---

### 11. Future Enhancements

- Persist data to a file or database.
- Add support for borrowing and returning books.
- Provide a web-based interface.

---

## Example Usage

### Add a New Book

```bash
./lms add "The Go Programming Language" "Alan A. A. Donovan" 2015 "978-0134190440"
```

### List All Books

```bash
./lms list
```

### View Book Details

```bash
./lms view 1
```

### Update a Book

```bash
./lms update 1 "The Go Programming Language" "Alan A. A. Donovan & Brian W. Kernighan" 2016 "978-0134190440"
```

### Delete a Book

```bash
./lms delete 1
```

---

This FSD provides a clear outline of what the Library Management System (LMS) application should do. You can use this document as a guide to implement the application in Go. Let me know if you'd like further assistance!
