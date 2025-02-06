package main

import "fmt"

// Book struct represents a book in the library.
type Book struct {
    ID     string
    Title  string
    Author string
}

// User struct represents a user of the library.
type User struct {
    ID         string
    Name       string
    BorrowedID string // ID of borrowed book; empty if none.
}

// Library struct represents the main library system.
type Library struct {
    Books []Book
    Users []User
}

// AddBook adds a new book to the library.
func (lib *Library) AddBook(book Book) {
    lib.Books = append(lib.Books, book)
    fmt.Printf("Book added: %s\n", book.Title)
}

// RemoveBook removes a book from the library.
func (lib *Library) RemoveBook(bookID string) error {
    for i, book := range lib.Books {
        if book.ID == bookID {
            lib.Books = append(lib.Books[:i], lib.Books[i+1:]...)
            fmt.Printf("Book removed: %s\n", book.Title)
            return nil
        }
    }
    return fmt.Errorf("Book not found")
}

// ListBooks lists all available books in the library.
func (lib *Library) ListBooks() {
    fmt.Println("Books in Library:")
    for _, book := range lib.Books {
        fmt.Printf("ID: %s, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
    }
}

// BorrowBook allows a user to borrow a book from the library.
func (lib *Library) BorrowBook(userID, bookID string) error {
    user, book := lib.findUser(userID), lib.findBook(bookID)

    if user == nil {
        return fmt.Errorf("User not found")
    }
    if book == nil {
        return fmt.Errorf("Book not found")
    }
    if user.BorrowedID != "" {
        return fmt.Errorf("User has already borrowed a book")
    }

    user.BorrowedID = bookID
    fmt.Printf("Book borrowed: %s by user %s\n", book.Title, user.Name)
    return nil
}

// ReturnBook allows a user to return a borrowed book to the library.
func (lib *Library) ReturnBook(userID, bookID string) error {
    user := lib.findUser(userID)

    if user == nil {
        return fmt.Errorf("User not found")
    }
    if user.BorrowedID != bookID {
        return fmt.Errorf("User did not borrow this book")
    }

    user.BorrowedID = ""
    fmt.Printf("Book returned: %s by user %s\n", bookID, user.Name)
    return nil
}

// ListUsers lists all users and their borrowed books.
func (lib *Library) ListUsers() {
    fmt.Println("Library Users:")
    for _, user := range lib.Users {
        borrowed := "None"
        if user.BorrowedID != "" {
            borrowed = user.BorrowedID
        }
        fmt.Printf("ID: %s, Name: %s, Borrowed Book ID: %s\n", user.ID, user.Name, borrowed)
    }
}

// findUser searches for a user by ID.
func (lib *Library) findUser(userID string) *User {
    for i := range lib.Users {
        if lib.Users[i].ID == userID {
            return &lib.Users[i]
        }
    }
    return nil
}

// findBook searches for a book by ID.
func (lib *Library) findBook(bookID string) *Book {
    for i := range lib.Books {
        if lib.Books[i].ID == bookID {
            return &lib.Books[i]
        }
    }
    return nil
}

// main function to run the library system example.
func main() {
    library := Library{}

    book1 := Book{ID: "1", Title: "Clean Code", Author: "Robert C. Martin"}
    user1 := User{ID: "1", Name: "John Doe"}

    library.AddBook(book1)
    library.Users = append(library.Users, user1)

    library.ListBooks()

    err := library.BorrowBook("1", "1")
    if err != nil {
        fmt.Println("Error:", err)
    }

    library.ListUsers()

    err = library.ReturnBook("1", "1")
    if err != nil {
        fmt.Println("Error:", err)
    }

    library.RemoveBook("1")

    library.ListBooks()
}
