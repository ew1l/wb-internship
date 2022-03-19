package unpack

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

// Unpack returns the unpacked string
func Unpack(s string) (us string, err error) {
	if _, err := strconv.Atoi(s); err == nil {
		return us, errors.New("некорректная строка")
	}

	var builder strings.Builder
	var previous rune
	var escaped bool

	for _, r := range s {
		if unicode.IsDigit(r) && !escaped {
			us = strings.Repeat(string(previous), int(r-'0')-1)
			builder.WriteString(us)
		} else {
			escaped = string(previous) != "\\" && string(r) == "\\"
			if !escaped {
				builder.WriteRune(r)
			}
			previous = r
		}
	}

	return builder.String(), nil
}
