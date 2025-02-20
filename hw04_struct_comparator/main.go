package main

import (
	"fmt"

	"github.com/Bladforceone/go_hw_otus/hw04_struct_comparator/comparator"
	"github.com/Bladforceone/go_hw_otus/hw04_struct_comparator/types"
)

func main() {
	book := types.Book{}
	book.SetTitle("kNigga")
	book.SetAuthor("Ryan Gosling")
	book.SetYear(2049)

	book2 := types.Book{}
	book2.SetYear(1917)

	comp := comparator.NewComparator(comparator.CompareByYear)

	fmt.Print(comp.Compare(book, book2))

	fmt.Printf("Book title: %s, Author: %s, Year: %d\n", book.Title(), book.Author(), book.Year())
}
