package printer

import (
	"fmt"

	"github.com/Bladforceone/go_hw_otus/hw06_testing/hw02/types"
)

func PrintStaff(staff []types.Employee) {
	for i := 0; i < len(staff); i++ {
		str := fmt.Sprintf("User ID: %d; Age: %d; Name: %s; Department ID: %d; ",
			staff[i].UserID, staff[i].Age, staff[i].Name, staff[i].DepartmentID)
		fmt.Println(str)
	}
}
