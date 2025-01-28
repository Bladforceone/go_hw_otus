package main

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSendGetRequest(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status": "ok"}`))
	}))
	defer ts.Close()

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	sendGetRequest(context.Background(), ts.URL)

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	io.Copy(&buf, r)

	assert.Contains(t, buf.String(), "Статус ответа:200 OK\nТело ответа(JSON):\n{\n  \"status\": \"ok\"\n}\n")
	assert.Contains(t, buf.String(), `"status": "ok"`)
}

func TestSendPostRequest(t *testing.T) {
	tests := []struct {
		name     string
		data     string
		expected string
	}{
		{
			name:     "Valid JSON",
			data:     `{"test": "ok"}`,
			expected: `"result": "success"`,
		},
		{
			name:     "Invalid JSON",
			data:     `invalid json`,
			expected: "Ошибка при парсинге JSON",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if tt.name == "Invalid JSON" {
					w.Write([]byte(`invalid json`))
				} else {
					w.Write([]byte(`{"result": "success"}`))
				}
			}))
			defer ts.Close()

			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			sendPostRequest(context.Background(), ts.URL, tt.data)

			w.Close()
			os.Stdout = oldStdout

			var buf bytes.Buffer
			io.Copy(&buf, r)

			assert.Contains(t, buf.String(), tt.expected)
		})
	}
}

func TestContextTimeout(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(100 * time.Millisecond)
		w.Write([]byte(`{}`))
	}))
	defer ts.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	sendGetRequest(ctx, ts.URL)

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	io.Copy(&buf, r)

	assert.Contains(t, buf.String(), "Ошибка при выполнении GET-запроса")
	assert.Contains(t, buf.String(), "context deadline exceeded")
}
