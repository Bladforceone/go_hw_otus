package book

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

func (b *Book) setID(id int) {
	b.id = id
}

func (b Book) getID() int {
	return b.id
}

func (b *Book) setTitle(title string) {
	b.title = title
}

func (b Book) getTitle() string {
	return b.title
}

func (b *Book) setAuthor(author string) {
	b.author = author
}

func (b Book) getAuthor() string {
	return b.author
}

func (b *Book) setYear(year int) {
	b.year = year
}

func (b Book) getYear() int {
	return b.year
}

func (b *Book) setSize(size int) {
	b.size = size
}

func (b Book) getSize() int {
	return b.size
}

func (b *Book) setRate(rate float32) {
	b.rate = rate
}

func (b Book) getRate() float32 {
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
