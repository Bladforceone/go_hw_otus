package types

type Book struct {
	id     int
	title  string
	author string
	year   int
	size   int
	rate   float32
}

func NewBook() *Book {
	b := new(Book)
	b.id = 0
	b.title = "empty"
	b.author = "empty"
	b.year = 0
	b.size = 0
	b.rate = 0.0
	return b
}

func (b *Book) CopyBook(other *Book) {
	b.id = other.id
	b.title = other.title
	b.author = other.author
	b.year = other.year
	b.size = other.size
	b.rate = other.rate
}

func (b *Book) ID() int {
	return b.id
}

func (b *Book) SetTitle(title string) {
	b.title = title
}

func (b *Book) Title() string {
	return b.title
}

func (b *Book) SetAuthor(author string) {
	b.author = author
}

func (b *Book) Author() string {
	return b.author
}

func (b *Book) SetYear(year int) {
	b.year = year
}

func (b *Book) Year() int {
	return b.year
}

func (b *Book) SetSize(size int) {
	b.size = size
}

func (b *Book) Size() int {
	return b.size
}

func (b *Book) SetRate(rate float32) {
	b.rate = rate
}

func (b *Book) Rate() float32 {
	return b.rate
}
