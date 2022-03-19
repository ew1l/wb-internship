package main

import (
	"fmt"
	"os"

	"github.com/ew1l/wb-l2/develop/dev09/wget"
)

func main() {
	w := wget.NewWget(1)

	if err := w.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
