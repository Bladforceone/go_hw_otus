package types

import "encoding/json"

type Book struct {
	ID     int
	Title  string
	Author string
	Year   int
	Size   int
	Rate   float64
}

func (b *Book) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ID     int     `json:"id"`
		Title  string  `json:"title"`
		Author string  `json:"author"`
		Year   int     `json:"year"`
		Size   int     `json:"size"`
		Rate   float64 `json:"rate"`
	}{
		ID:     b.ID,
		Title:  b.Title,
		Author: b.Author,
		Year:   b.Year,
		Size:   b.Size,
		Rate:   b.Rate,
	})
}

func (b *Book) UnmarshalJSON(data []byte) error {
	var aux struct {
		ID     int     `json:"id"`
		Title  string  `json:"title"`
		Author string  `json:"author"`
		Year   int     `json:"year"`
		Size   int     `json:"size"`
		Rate   float64 `json:"rate"`
	}
	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}
	b.ID = aux.ID
	b.Title = aux.Title
	b.Author = aux.Author
	b.Year = aux.Year
	b.Size = aux.Size
	b.Rate = aux.Rate
	return nil
}
