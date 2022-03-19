package anagram

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	tests := []struct {
		words    []string
		expected map[string][]string
	}{
		{
			words:    []string{"пятак", "листок", "листок", "пяТка", "привет", "ТЯПКА", "столик", "слиТок", "тевирп"},
			expected: map[string][]string{"листок": {"слиток", "столик"}, "пятак": {"пятка", "тяпка"}},
		},
		{
			words:    []string{"актёр", "сетКа", "тёрка", "рОГа", "сетка", "гора", "тесАК", "аскет", "арго"},
			expected: map[string][]string{"рога": {"арго", "гора"}, "сетка": {"аскет", "тесак"}},
		},
	}

	for name, test := range tests {
		t.Run(strconv.Itoa(name), func(t *testing.T) {
			actual := Search(test.words)
			assert.Equal(t, test.expected, actual)
		})
	}
}
