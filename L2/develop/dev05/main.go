package main

import (
	"fmt"
	"os"

	"github.com/ew1l/wb-l2/develop/dev05/grep"
)

func main() {
	g := grep.NewGrep()

	if err := g.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
