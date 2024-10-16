package main

import (
	"fmt"

	"github.com/Bladforceone/go_hw_otus/hw02_fix_app/printer"
	"github.com/Bladforceone/go_hw_otus/hw02_fix_app/reader"
	"github.com/Bladforceone/go_hw_otus/hw02_fix_app/types"
)

func main() {
	var path string

	fmt.Printf("Enter data file path: ")
	fmt.Scanln(&path)

	if len(path) == 0 {
		path = "data.json"
	}

	var err error
	var staff []types.Employee

	staff, err = reader.ReadJSON(path)
	if err != nil {
		fmt.Print(err)
	}

	printer.PrintStaff(staff)
}
