package types

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"encoding/xml"

	"go.mongodb.org/mongo-driver/bson"
	"gopkg.in/yaml.v3"
)

type Book struct {
	ID     int     `json:"id" xml:"id" yaml:"id"`
	Title  string  `json:"title" xml:"title" yaml:"title"`
	Author string  `json:"author" xml:"author" yaml:"author"`
	Year   int     `json:"year" xml:"year" yaml:"year"`
	Size   int     `json:"size" xml:"size" yaml:"size"`
	Rate   float32 `json:"rate" xml:"rate" yaml:"rate"`
	Sample []byte  `json:"sample" xml:"sample" yaml:"sample"`
}

func UnmarshalJSONSlices(data []byte) ([]Book, error) {
	var tmp []Book
	return tmp, json.Unmarshal(data, &tmp)
}

func MarshalJSONSlices(books []Book) ([]byte, error) {
	return json.Marshal(books)
}

func UnmarshalXMLSlices(data []byte) ([]Book, error) {
	var tmp []Book
	return tmp, xml.Unmarshal(data, &tmp)
}

func MarshalXMLSlices(books []Book) ([]byte, error) {
	return xml.Marshal(books)
}

func UnmarshalYAMLSlices(data []byte) ([]Book, error) {
	var tmp []Book
	return tmp, yaml.Unmarshal(data, &tmp)
}

func MarshalYAMLSlices(books Book) ([]byte, error) {
	return yaml.Marshal(books)
}

func (b *Book) UnmarshalGob(data []byte) error {
	dec := gob.NewDecoder(bytes.NewReader(data))
	if err := dec.Decode(b); err != nil {
		return err
	}
	return nil
}

func (b *Book) MarshalGob() ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(b); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (b *Book) UnmarshalBSON(data []byte) error {
	var tmp Book
	if err := bson.Unmarshal(data, &tmp); err != nil {
		return err
	}
	return nil
}

func (b *Book) MarshalBSON() ([]byte, error) {
	return bson.Marshal(b)
}
