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

type Comparator struct {
	Mode CompareMode
}

func NewComparator(mod CompareMode) Comparator {
	c := Comparator{}
	c.Mode = mod
	return c
}

func (c *Comparator) Compare(b1, b2 types.Book) (bool, error) {
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
