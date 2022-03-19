package unpack

import (
	"errors"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnpack(t *testing.T) {
	type args struct {
		s   string
		err error
	}

	tests := []struct {
		s        string
		expected args
	}{
		{
			s: "a4bc2d5e",
			expected: args{
				s:   "aaaabccddddde",
				err: nil,
			},
		},
		{
			s: "abcd",
			expected: args{
				s:   "abcd",
				err: nil,
			},
		},
		{
			s: "45",
			expected: args{
				s:   "",
				err: errors.New("некорректная строка"),
			},
		},
		{
			s: "",
			expected: args{
				s:   "",
				err: nil,
			},
		},
		{
			s: `qwe\4\5`,
			expected: args{
				s:   "qwe45",
				err: nil,
			},
		},
		{
			s: `qwe\45`,
			expected: args{
				s:   "qwe44444",
				err: nil,
			},
		},
		{
			s: `qwe\\5`,
			expected: args{
				s:   `qwe\\\\\`,
				err: nil,
			},
		},
		{
			s: `прив\4ет`,
			expected: args{
				s:   "прив4ет",
				err: nil,
			},
		},
	}

	for name, test := range tests {
		t.Run(strconv.Itoa(name), func(t *testing.T) {
			actual, err := Unpack(test.s)
			assert.Equal(t, test.expected, args{actual, err})
		})
	}
}
