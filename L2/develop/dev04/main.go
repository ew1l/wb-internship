package main

import (
	"fmt"

	"L2/develop/dev04/anagram"
)

func main() {
	words := []string{
		"пятак",
		"листок",
		"листок",
		"пяТка",
		"привет",
		"ТЯПКА",
		"столик",
		"слиТок",
		"тевирп",
	}

	setas := anagram.Search(words)
	fmt.Println(setas)
}

// map[листок:[слиток столик] пятак:[пятка тяпка]]
