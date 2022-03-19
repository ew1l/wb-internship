package anagram

import (
	"sort"
	"strings"
)

// Search returns a set of anagrams
func Search(words []string) map[string][]string {
	setas := make(map[string][]string)
	temp := make(map[string]string)

	for _, word := range words {
		word := strings.ToLower(word)

		sword := func(word string) string {
			r := []rune(word)

			sort.Slice(r, func(i, j int) bool {
				return r[i] < r[j]
			})

			return string(r)
		}(word)

		if value, ok := temp[sword]; ok && value != word {
			setas[value] = append(setas[value], word)
		} else {
			temp[sword] = word
		}
	}

	for key, seta := range setas {
		if len(seta) < 2 {
			delete(setas, key)
		}
		sort.Strings(seta)
	}

	return setas
}
