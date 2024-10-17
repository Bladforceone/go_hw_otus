package main

import (
	"fmt"

	"github.com/Bladforceone/go_hw_otus/hw04_struct_comparator/types"
)

func main() {
	book := types.Book{}
	book.SetTitle("kNigga")
	book.SetAuthor("Ryan Gosling")
	book.SetYear(2049)

	book2 := book
	book2.SetYear(1917)
	fmt.Printf("Book title: %s, Author: %s, Year: %d\n", book.GetTitle(), book.GetAuthor(), book.GetYear())
}
