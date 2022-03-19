package main

import (
	"fmt"

	"L2/develop/dev03/sort"
)

func main() {
	s := sort.NewSort()

	if err := s.Execute(); err != nil {
		fmt.Println(err)
	}
}
