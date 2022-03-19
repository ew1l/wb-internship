package main

import (
	"fmt"

	"github.com/ew1l/wb-l2/develop/dev06/cut"
)

func main() {
	c := cut.NewCut()

	if err := c.Execute(); err != nil {
		fmt.Println(err)
	}
}
