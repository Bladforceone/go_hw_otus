package reader

import (
	"encoding/json"
	"io"
	"os"

	"github.com/Bladforceone/go_hw_otus/hw06_testing/hw02/types"
)

func ReadJSON(filePath string) ([]types.Employee, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	bytes, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	var data []types.Employee

	err = json.Unmarshal(bytes, &data)

	return data, err
}
