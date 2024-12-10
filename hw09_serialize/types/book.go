package types

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"encoding/xml"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"gopkg.in/yaml.v3"
)

type Book struct {
	ID     int     `json:"id,omitempty" xml:"id,omitempty" yaml:"id,omitempty" bson:"id,omitempty"`
	Title  string  `json:"title,omitempty" xml:"title,omitempty" yaml:"title,omitempty" bson:"title"`
	Author string  `json:"author,omitempty" xml:"author,omitempty" yaml:"author,omitempty" bson:"author"`
	Year   int     `json:"year,omitempty" xml:"year,omitempty" yaml:"year,omitempty" bson:"year,omitempty"`
	Size   int     `json:"size,omitempty" xml:"size,omitempty" yaml:"size,omitempty" bson:"size,omitempty"`
	Rate   float32 `json:"rate,omitempty" xml:"rate,omitempty" yaml:"rate,omitempty" bson:"rate,omitempty"`
	Sample []byte  `json:"sample,omitempty" xml:"sample,omitempty" yaml:"sample,omitempty" bson:"sample,omitempty"`
}

func UnmarshalJSONSlices(data []byte) ([]Book, error) {
	var books []Book
	err := json.Unmarshal(data, &books)
	return books, err
}

func MarshalJSONSlices(books []Book) ([]byte, error) {
	return json.Marshal(books)
}

func UnmarshalXMLSlices(data []byte) ([]Book, error) {
	var booksWrapper struct {
		Books []Book `xml:"book"`
	}
	err := xml.Unmarshal(data, &booksWrapper)
	return booksWrapper.Books, err
}

func MarshalXMLSlices(books []Book) ([]byte, error) {
	booksWrapper := struct {
		XMLName xml.Name `xml:"books"`
		Books   []Book   `xml:"book"`
	}{
		Books: books,
	}
	return xml.Marshal(booksWrapper)
}

func UnmarshalYAMLSlices(data []byte) ([]Book, error) {
	var books []Book
	err := yaml.Unmarshal(data, &books)
	return books, err
}

func MarshalYAMLSlices(books []Book) ([]byte, error) {
	return yaml.Marshal(books)
}

func DeserializeFromGOB(data []byte, v interface{}) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	err := dec.Decode(v)
	if err != nil {
		return fmt.Errorf("failed to deserialize: %w", err)
	}
	return nil
}
func SerializeToGOB(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(v)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize: %w", err)
	}
	return buf.Bytes(), nil
}

func SerializeToBSON(v interface{}) ([]byte, error) {
	data, err := bson.Marshal(v)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize: %w", err)
	}
	return data, nil
}

func DeserializeFromBSON(data []byte, v interface{}) error {
	err := bson.Unmarshal(data, v)
	if err != nil {
		return fmt.Errorf("failed to deserialize: %w", err)
	}
	return nil
}
