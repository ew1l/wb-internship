package sort

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

type flags struct {
	k *int
	n *bool
	r *bool
	u *bool
}

// Sort structure
type Sort struct {
	flags
	data []string
}

// NewSort sort's structure constructor
func NewSort() *Sort {
	return &Sort{
		flags: flags{
			k: flag.Int("k", 1, "specifying a column to sort"),
			n: flag.Bool("n", false, "sort by numeric value"),
			r: flag.Bool("r", false, "sort in reverse order"),
			u: flag.Bool("u", false, "do not output duplicate lines"),
		},
		data: make([]string, 0),
	}
}

// Execute runs the utility
func (s *Sort) Execute() error {
	flag.Parse()

	data, err := ioutil.ReadFile(os.Args[len(os.Args)-1])
	if err != nil {
		return err
	}

	s.data = strings.Split(strings.TrimSpace(string(data)), "\n")

	sort.Slice(s.data, func(i, j int) bool {
		return strings.Split(s.data[i], " ")[*s.flags.k-1] < strings.Split(s.data[j], " ")[*s.flags.k-1]
	})

	if *s.flags.n {
		sort.Slice(s.data, func(i, j int) bool {
			fline, _ := strconv.Atoi(strings.Split(s.data[i], " ")[*s.flags.k-1])
			sline, _ := strconv.Atoi(strings.Split(s.data[j], " ")[*s.flags.k-1])

			return fline < sline
		})
	}

	if *s.flags.r {
		for i, j := 0, len(s.data)-1; i < j; i, j = i+1, j-1 {
			s.data[i], s.data[j] = s.data[j], s.data[i]
		}
	}

	if *s.flags.u {
		temp := make(map[string]struct{})
		d := make([]string, 0)

		for _, line := range s.data {
			if _, ok := temp[line]; !ok {
				temp[line] = struct{}{}
				d = append(d, line)
			}
		}

		s.data = d
	}

	for _, line := range s.data {
		fmt.Println(line)
	}

	return nil
}
