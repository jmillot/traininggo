package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var schema = `
CREATE TABLE IF NOT EXISTS book
(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	title TEXT,
	author TEXT,
	page_count INTEGER
)
`

type Book struct {
	Id        int
	Title     string
	Author    string
	PageCount int
}

func main() {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// insert INTO
	book := Book{
		Title:     "Harry Potter",
		Author:    "J.K. Rowling",
		PageCount: 333,
	}

	stmt, _ := db.Prepare("INSERT INTO book (title, author, page_count) VALUES (?, ?, ?)")
	_, err = stmt.Exec(book.Title, book.Author, book.PageCount)
	if err != nil {
		log.Fatal(err)
	}

	stmt, err = db.Prepare(schema)
	if err != nil {
		log.Fatal(err)
	}
	stmt.Exec()

	// map to values
	var id int
	var title string
	var author string
	var pageCount int

	rows, _ := db.Query("SELECT * FROM book")
	for rows.Next() {
		rows.Scan(&id, &title, &author, &pageCount)
		fmt.Printf("id=%v, title=%v, author=%v, pageCount=%v\n", id, title, author, pageCount)
	}

	// Map to Struct
	rows, _ = db.Query("SELECT * FROM book")
	b := Book{}
	for rows.Next() {
		rows.Scan(b.Id, b.Title, b.Author, b.PageCount)
		fmt.Printf("id=%v, title=%v, author=%v, pageCount=%v\n", b.Id, b.Title, b.Author, b.PageCount)
	}

	fmt.Println("Excecute to database successfull !")
}
