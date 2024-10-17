package types

import "fmt"

type CompareMode int

const (
	CompareByYear CompareMode = iota
	CompareBySize
	CompareByRate
)

type Book struct {
	id     int
	title  string
	author string
	year   int
	size   int
	rate   float32
}

func (b *Book) SetID(id int) {
	b.id = id
}

func (b Book) GetID() int {
	return b.id
}

func (b *Book) SetTitle(title string) {
	b.title = title
}

func (b Book) GetTitle() string {
	return b.title
}

func (b *Book) SetAuthor(author string) {
	b.author = author
}

func (b Book) GetAuthor() string {
	return b.author
}

func (b *Book) SetYear(year int) {
	b.year = year
}

func (b Book) GetYear() int {
	return b.year
}

func (b *Book) SetSize(size int) {
	b.size = size
}

func (b Book) GetSize() int {
	return b.size
}

func (b *Book) SetRate(rate float32) {
	b.rate = rate
}

func (b Book) GetRate() float32 {
	return b.rate
}

func (b Book) Compare(other Book, mod CompareMode) (bool, error) {
	switch mod {
	case CompareByRate:
		return b.rate > other.rate, nil
	case CompareBySize:
		return b.size > other.size, nil
	case CompareByYear:
		return b.year > other.year, nil
	default:
		return false, fmt.Errorf("uncorrected compare mode")
	}
}
