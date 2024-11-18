package counter_test

import (
	"testing"

	"github.com/Bladforceone/go_hw_otus/hw07_word_counter/counter"
	"github.com/stretchr/testify/assert"
)

func TestCounterWords(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want map[string]int
	}{
		{name: "empty string", str: "", want: nil},
		{name: "single word", str: "hello", want: map[string]int{"hello": 1}},
		{name: "multiple words", str: "hello world", want: map[string]int{"hello": 1, "world": 1}},
		{name: "input with symbol", str: "hello world!", want: map[string]int{"hello": 1, "world": 1}},
		{name: "input with capital letters", str: "Hello World!", want: map[string]int{"hello": 1, "world": 1}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := counter.CountWords(test.str)
			assert.Equal(t, test.want, got)
		})
	}
}
