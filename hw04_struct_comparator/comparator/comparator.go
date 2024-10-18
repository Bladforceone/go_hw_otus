package comparator

import (
	"fmt"

	"github.com/Bladforceone/go_hw_otus/hw04_struct_comparator/types"
)

type CompareMode int

const (
	CompareByYear CompareMode = iota
	CompareBySize
	CompareByRate
)

type comparator struct {
	Mode CompareMode
}

func NewComparator(mod CompareMode) *comparator {
	c := new(comparator)
	c.Mode = mod
	return c
}

func (c *comparator) Compare(b1, b2 types.Book) (bool, error) {
	switch c.Mode {
	case CompareByRate:
		return b1.Rate() > b2.Rate(), nil
	case CompareBySize:
		return b1.Size() > b2.Size(), nil
	case CompareByYear:
		return b1.Year() > b2.Year(), nil
	default:
		return false, fmt.Errorf("uncorrected compare mode")
	}
}
