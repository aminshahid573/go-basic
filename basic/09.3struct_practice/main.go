package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Pages  int
}

func (b Book) PrintDetails() {
	fmt.Printf("Book: %s by %s (%d pages)\n", b.Title, b.Author, b.Pages)
}

func main() {
	book1 := Book{"The Alchemist", "Paulo Coelho", 208}
	book2 := Book{"The Lord of the Rings", "J.R.R Tolkien", 1216}

	book1.PrintDetails()
	book2.PrintDetails()

}
