// --Summary:
//
//	Create a program to manage lending of library books.
//
// --Requirements:
// * The library must have books and members, and must include:
//   - Which books have been checked out
//   - What time the books were checked out
//   - What time the books were returned
//
// * Perform the following:
//   - Add at least 4 books and at least 3 members to the library
//   - Check out a book
//   - Check in a book
//   - Print out initial library information, and after each change
//
// * There must only ever be one copy of the library in memory at any time
//
// --Notes:
// * Use the `time` package from the standard library for check in/out times
// * Liberal use of type aliases, structs, and maps will help organize this project
package main

import (
	"fmt"
	"time"
)

type Book struct {
	ID             int
	Title          string
	Author         string
	CheckedOut     bool
	CheckedOutTime time.Time
	ReturnedTime   time.Time
}

type Member struct {
	ID   int
	Name string
}

type Library struct {
	Books   map[int]Book
	Members map[int]Member
}

func (lib *Library) AddBook(book Book) {
	lib.Books[book.ID] = book
}

func (lib *Library) AddMember(member Member) {
	lib.Members[member.ID] = member
}

func (lib *Library) CheckOutBook(bookID, memberID int) {
	book, ok := lib.Books[bookID]
	if !ok {
		fmt.Println("Book not found")
		return
	}
	if book.CheckedOut {
		fmt.Println("Book already checked out")
		return
	}

	member, ok := lib.Members[memberID]
	if !ok {
		fmt.Println("Member not found")
		return
	}

	lib.Books[bookID] = Book{
		ID:             book.ID,
		Title:          book.Title,
		Author:         book.Author,
		CheckedOut:     true,
		CheckedOutTime: time.Now(),
	}

	fmt.Printf("Book '%s' checked out by %s\n", book.Title, member.Name)
}

func (lib *Library) CheckInBook(bookID int) {
	book, ok := lib.Books[bookID]
	if !ok {
		fmt.Println("Book not found")
		return
	}
	if !book.CheckedOut {
		fmt.Println("Book already checked in")
		return
	}

	lib.Books[bookID] = Book{
		ID:           book.ID,
		Title:        book.Title,
		Author:       book.Author,
		CheckedOut:   false,
		ReturnedTime: time.Now(),
	}

	fmt.Printf("Book '%s' checked in\n", book.Title)
}

func (lib *Library) PrintLibraryInfo() {
	fmt.Println("\nLibrary Information:")
	fmt.Println("Books:")
	for _, book := range lib.Books {
		fmt.Printf("- %s by %s\n", book.Title, book.Author)
		if book.CheckedOut {
			fmt.Printf("  Checked Out: Yes (Time: %s)\n", book.CheckedOutTime.Format("2006-01-02 15:04:05"))
		} else {
			fmt.Println("  Checked Out: No")
		}
	}
	fmt.Println("\nMembers:")
	for _, member := range lib.Members {
		fmt.Printf("- %s\n", member.Name)
	}
	fmt.Println()
}

func main() {
	lib := Library{
		Books:   make(map[int]Book),
		Members: make(map[int]Member),
	}

	// Add books
	lib.AddBook(Book{ID: 1, Title: "Go BootCamp", Author: "Techie Emma"})
	lib.AddBook(Book{ID: 2, Title: "Programming In Go", Author: "Zero To Mastery"})
	lib.AddBook(Book{ID: 3, Title: "Api Security Testing", Author: "Said Lassri"})
	lib.AddBook(Book{ID: 4, Title: "Go Developer Guide", Author: "Jayson Doe"})

	// Add members
	lib.AddMember(Member{ID: 1, Name: "Emmanuel Nicholas"})
	lib.AddMember(Member{ID: 2, Name: "Said Developer"})
	lib.AddMember(Member{ID: 3, Name: "Oyinlade Ojasenmi"})

	// Print initial library information
	lib.PrintLibraryInfo()

	// Check out a book
	lib.CheckOutBook(1, 1)
	lib.CheckOutBook(2, 3)
	lib.PrintLibraryInfo()

	// Check in a book
	lib.CheckInBook(1)
	lib.PrintLibraryInfo()
}
