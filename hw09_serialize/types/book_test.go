package types_test

import (
	"testing"

	"github.com/Bladforceone/go_hw_otus/hw09_serialize/types"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestUnmarshalJSONSlices(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		want    []types.Book
		wantErr bool
	}{
		{
			name: "valid JSON data",
			data: []byte(`[
                {"title": "Книга 1", "author": "Автор 1"},
                {"title": "Книга 2", "author": "Автор 2"}
            ]`),
			want: []types.Book{
				{Title: "Книга 1", Author: "Автор 1"},
				{Title: "Книга 2", Author: "Автор 2"},
			},
			wantErr: false,
		},
		{
			name:    "invalid JSON data",
			data:    []byte(`{"le": "Книга 1", "aor": "Автор 1"}`),
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := types.UnmarshalJSONSlices(tt.data)
			if !tt.wantErr {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			} else {
				assert.Error(t, err)
				assert.Nil(t, got)
			}
		})
	}
}

func TestMarshalJSONSlices(t *testing.T) {
	tests := []struct {
		name    string
		data    []types.Book
		want    []byte
		wantErr bool
	}{
		{
			name: "valid JSON data",
			data: []types.Book{
				{Title: "Книга 1", Author: "Автор 1"},
				{Title: "Книга 2", Author: "Автор 2"},
			},
			want: []byte(`[
                {"title":"Книга 1","author":"Автор 1"},
                {"title":"Книга 2","author":"Автор 2"}
            ]`),
			wantErr: false,
		},
		{
			name:    "empty data",
			data:    []types.Book{},
			want:    []byte("[]"),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := types.MarshalJSONSlices(tt.data)
			if !tt.wantErr {
				assert.NoError(t, err)
				assert.JSONEq(t, string(tt.want), string(got))
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestUnmarshalXMLSlices(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		want    []types.Book
		wantErr bool
	}{
		{
			name: "valid XML data",
			data: []byte(`
                <books>
                    <book>
                        <title>Книга 1</title>
                        <author>Автор 1</author>
                    </book>
                    <book>
                        <title>Книга 2</title>
                        <author>Автор 2</author>
                    </book>
                </books>
            `),
			want: []types.Book{
				{Title: "Книга 1", Author: "Автор 1"},
				{Title: "Книга 2", Author: "Автор 2"},
			},
			wantErr: false,
		},
		{
			name:    "invalid XML data",
			data:    []byte(`rht<invalid>tytrh</inv`),
			want:    nil,
			wantErr: true,
		},
		{
			name:    "empty XML data",
			data:    []byte(""),
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := types.UnmarshalXMLSlices(tt.data)
			if !tt.wantErr {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			} else {
				assert.Error(t, err)
				assert.Nil(t, got)
			}
		})
	}
}

func TestMarshalXMLSlices(t *testing.T) {
	tests := []struct {
		name    string
		data    []types.Book
		want    []byte
		wantErr bool
	}{
		{
			name: "valid XML data",
			data: []types.Book{
				{Title: "Книга 1", Author: "Автор 1"},
				{Title: "Книга 2", Author: "Автор 2"},
			},
			want: []byte{
				0x3c, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x3e, 0x3c, 0x62, 0x6f, 0x6f, 0x6b, 0x3e, 0x3c, 0x74, 0x69,
				0x74, 0x6c, 0x65, 0x3e, 0xd0, 0x9a, 0xd0, 0xbd, 0xd0, 0xb8, 0xd0, 0xb3, 0xd0, 0xb0, 0x20, 0x31, 0x3c,
				0x2f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x3e, 0x3c, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x3e, 0xd0, 0x90,
				0xd0, 0xb2, 0xd1, 0x82, 0xd0, 0xbe, 0xd1, 0x80, 0x20, 0x31, 0x3c, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f,
				0x72, 0x3e, 0x3c, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x3e, 0x3c, 0x62, 0x6f, 0x6f, 0x6b, 0x3e, 0x3c, 0x74,
				0x69, 0x74, 0x6c, 0x65, 0x3e, 0xd0, 0x9a, 0xd0, 0xbd, 0xd0, 0xb8, 0xd0, 0xb3, 0xd0, 0xb0, 0x20, 0x32,
				0x3c, 0x2f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x3e, 0x3c, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x3e, 0xd0,
				0x90, 0xd0, 0xb2, 0xd1, 0x82, 0xd0, 0xbe, 0xd1, 0x80, 0x20, 0x32, 0x3c, 0x2f, 0x61, 0x75, 0x74, 0x68,
				0x6f, 0x72, 0x3e, 0x3c, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x3e, 0x3c, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x73,
				0x3e,
			},
			wantErr: false,
		},
		{
			name:    "empty data",
			data:    []types.Book{},
			want:    []byte("<books></books>"),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := types.MarshalXMLSlices(tt.data)
			if !tt.wantErr {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestUnmarshalYAMLSlices(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		want    []types.Book
		wantErr bool
	}{
		{
			name: "valid YAML data",
			data: []byte(`
- title: "Книга 1"
  author: "Автор 1"
- title: "Книга 2"
  author: "Автор 2"
`),
			want: []types.Book{
				{Title: "Книга 1", Author: "Автор 1"},
				{Title: "Книга 2", Author: "Автор 2"},
			},
			wantErr: false,
		},
		{
			name:    "invalid YAML data",
			data:    []byte(`- title: "Книга 1" author: "Автор 1"`),
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := types.UnmarshalYAMLSlices(tt.data)
			if !tt.wantErr {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			} else {
				assert.Error(t, err)
				assert.Nil(t, got)
			}
		})
	}
}

func TestMarshalYAMLSlices(t *testing.T) {
	tests := []struct {
		name    string
		data    []types.Book
		want    []byte
		wantErr bool
	}{
		{
			name: "valid YAML data",
			data: []types.Book{
				{Title: "Книга 1", Author: "Автор 1"},
				{Title: "Книга 2", Author: "Автор 2"},
			},
			want: []byte(`
- title: "Книга 1"
  author: "Автор 1"
- title: "Книга 2"
  author: "Автор 2"
`),
			wantErr: false,
		},
		{
			name:    "empty data",
			data:    []types.Book{},
			want:    []byte("[]interface {}{}"),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := types.MarshalYAMLSlices(tt.data)
			if !tt.wantErr {
				assert.NoError(t, err)
				assert.YAMLEq(t, string(tt.want), string(got))
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestSerializeToGOB(t *testing.T) {
	tests := []struct {
		name    string
		input   *types.Book
		wantErr bool
	}{
		{
			name: "valid data",
			input: &types.Book{
				ID:     1,
				Title:  "Go Programming",
				Author: "John Doe",
				Year:   2024,
				Size:   300,
				Rate:   4.8,
				Sample: []byte("Sample data for Go Programming Book"),
			},
			wantErr: false,
		},
		{
			name: "empty data",
			input: &types.Book{
				ID:     0,
				Title:  "",
				Author: "",
				Year:   0,
				Size:   0,
				Rate:   0.0,
				Sample: nil,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serializedData, err := types.SerializeToGOB(tt.input)
			if !tt.wantErr {
				assert.NoError(t, err)
				assert.NotNil(t, serializedData)
				assert.Greater(t, len(serializedData), 0)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestDeserializeFromGOB(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		want    *types.Book
		wantErr bool
	}{
		{
			name: "valid data",
			data: func() []byte {
				// Подготавливаем данные для десериализации
				serializedData, _ := types.SerializeToGOB(&types.Book{
					ID:     1,
					Title:  "Go Programming",
					Author: "John Doe",
					Year:   2024,
					Size:   300,
					Rate:   4.8,
					Sample: []byte("Sample data for Go Programming Book"),
				})
				return serializedData
			}(),
			want: &types.Book{
				ID:     1,
				Title:  "Go Programming",
				Author: "John Doe",
				Year:   2024,
				Size:   300,
				Rate:   4.8,
				Sample: []byte("Sample data for Go Programming Book"),
			},
			wantErr: false,
		},
		{
			name:    "invalid data",
			data:    []byte(`invalid-gob-data`), // Невалидные данные GOB
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var deserializedBook types.Book
			err := types.DeserializeFromGOB(tt.data, &deserializedBook)
			if !tt.wantErr {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, &deserializedBook)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestDeserializeFromBSON(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		want    *types.Book
		wantErr bool
	}{
		{
			name: "valid data",
			data: func() []byte {
				data, _ := bson.Marshal(&types.Book{
					ID:     1,
					Title:  "The Go Programming Language",
					Author: "Alan A. A. Donovan",
					Year:   2015,
					Size:   400,
					Rate:   4.5,
					Sample: []byte("Sample text for Go book"),
				})
				return data
			}(),
			want: &types.Book{
				ID:     1,
				Title:  "The Go Programming Language",
				Author: "Alan A. A. Donovan",
				Year:   2015,
				Size:   400,
				Rate:   4.5,
				Sample: []byte("Sample text for Go book"),
			},
			wantErr: false,
		},
		{
			name:    "invalid BSON data",
			data:    []byte(`invalid-bson-data`),
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var deserializedBook types.Book
			err := types.DeserializeFromBSON(tt.data, &deserializedBook)
			if !tt.wantErr {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, &deserializedBook)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func SerializeToBSON(v interface{}) ([]byte, error) {
	return bson.Marshal(v)
}

// Функция десериализации из BSON
func DeserializeFromBSON(data []byte, v interface{}) error {
	return bson.Unmarshal(data, v)
}

func TestSerializeToBSON(t *testing.T) {
	tests := []struct {
		name    string
		input   *types.Book
		want    []byte
		wantErr bool
	}{
		{
			name: "valid data",
			input: &types.Book{
				ID:     1,
				Title:  "The Go Programming Language",
				Author: "Alan A. A. Donovan",
				Year:   2015,
				Size:   400,
				Rate:   4.5,
				Sample: []byte("Sample text for Go book"),
			},
			want: func() []byte {
				data, _ := bson.Marshal(&types.Book{
					ID:     1,
					Title:  "The Go Programming Language",
					Author: "Alan A. A. Donovan",
					Year:   2015,
					Size:   400,
					Rate:   4.5,
					Sample: []byte("Sample text for Go book"),
				})
				return data
			}(),
			wantErr: false,
		},
		{
			name: "invalid data",
			input: &types.Book{
				ID:     0,
				Title:  "",
				Author: "",
				Year:   0,
				Size:   0,
				Rate:   0.0,
				Sample: nil,
			},
			want: []byte{0x1e, 0x0, 0x0, 0x0, 0x2, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x0, 0x1,
				0x0, 0x0, 0x0, 0x0, 0x2, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x0, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serializedData, err := SerializeToBSON(tt.input)
			if !tt.wantErr {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, serializedData)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
