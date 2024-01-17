# Booky API Project

This project is a microservice-based REST API for managing a book library. It provides functionalities to add books, borrow books, and mark borrowed books as returned. The API is built using the Go programming language and the Gin web framework.

## Technologies Used

- [Go Programming Language](https://golang.org/): Go is used as the primary programming language for building the API.
- [Gin Web Framework](https://github.com/gin-gonic/gin): Gin is used as the web framework to create RESTful APIs.
- [GORM](https://gorm.io/): GORM is utilized as the Object Relational Mapping (ORM) library for interacting with the MySQL database.

## Prerequisites

Before running the application locally, ensure that you have the following installed:

- [Go](https://golang.org/dl/): The Go programming language.
- [Git](https://git-scm.com/): Version control system.

## Setup

1. Clone this repository to your local machine:

   ```sh
   git clone https://github.com/your-username/booky-api.git
   cd booky-api
   ```

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

### Get All Borrows

- **URL:** `/api/v1/booky/borrows`
- **Method:** `GET`
- **Description:** Fetches a list of all borrow records.

### Return a Borrowed Book

- **URL:** `/api/v1/booky/:book_id/borrow/:borrow_id`
- **Method:** `PUT`
- **Description:** Marks a borrowed book as returned by providing the book ID and borrow ID.

## Database Configuration

The application uses an online MySQL database provided by Aiven. You can find the database connection details in the `main.go` file. If needed, update the connection string with your own database credentials.

## Workflow

The API logic is organized into distinct segments:

- Adding a book to the library.
- Borrowing a book with validations for availability and borrowing duration.
- Returning a borrowed book.

Each logic segment in the code is commented for better understanding.

Feel free to contribute or extend this project as needed!
