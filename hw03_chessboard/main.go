package main

import (
	"fmt"
)

func printChessboard(size int) error {
	var evenRow, oddRow string

	if size < 2 {
		return fmt.Errorf("uncorrected size")
	}

	for i := 0; i < size; i++ {
		if i%2 == 0 {
			evenRow += " "
			oddRow += "#"
		} else {
			evenRow += "#"
			oddRow += " "
		}
	}

	for i := 0; i < size; i++ {
		if i%2 == 0 {
			fmt.Printf("%s \n", evenRow)
		} else {
			fmt.Printf("%s \n", oddRow)
		}
	}

	return nil
}

func main() {
	var size int

	fmt.Scan(&size)

	err := printChessboard(size)
	if err != nil {
		fmt.Print(err)
	}
}
