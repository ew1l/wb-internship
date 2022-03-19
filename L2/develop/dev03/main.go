package main

import (
	"fmt"

	"github.com/ew1l/wb-l2/develop/dev03/sort"
)

func main() {
	s := sort.NewSort()

	if err := s.Execute(); err != nil {
		fmt.Println(err)
	}
}
