# Booky API Project

This project is a simple REST API for managing a book library. It provides functionalities to add books, borrow books, and mark borrowed books as returned.

## Technologies Used

- [Go Programming Language](https://golang.org/): Go is used as the primary programming language for building the API.
- [Gin Web Framework](https://github.com/gin-gonic/gin): Gin is used as the web framework to create RESTful APIs.

## How to Run Locally

### Prerequisites

- Make sure you have Go installed. If not, you can download and install it from [Go's official website](https://golang.org/).
- Clone this repository to your local machine.

### Running the API

1. Navigate to the project directory.
2. Run the following command to start the server:

   ```sh
   go run main.go
   ```

3. The server will start on `localhost:8080`.

## API Endpoints

### Get All Books

- **URL:** `/api/v1/booky`
- **Method:** `GET`
- **Description:** Fetches a list of all books.

### Add a Book

- **URL:** `/api/v1/booky`
- **Method:** `POST`
- **Description:** Adds a new book to the library.

### Borrow a Book

- **URL:** `/api/v1/booky/:book_id/borrow`
- **Method:** `POST`
- **Description:** Borrows a book by specifying the book ID.

### Return a Borrowed Book

- **URL:** `/api/v1/booky/:book_id/borrow/:borrow_id`
- **Method:** `PUT`
- **Description:** Marks a borrowed book as returned by providing the book ID and borrow ID.

## Logic Overview

The API uses Go with the Gin web framework to manage book data. It includes logic to handle various scenarios:
- Adding a book to the library.
- Borrowing a book with validations for availability and borrowing duration.
- Returning a borrowed book.

Each logic segment in the code is commented for better understanding.

Feel free to contribute or extend this project as needed!
