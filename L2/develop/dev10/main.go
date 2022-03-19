package main

import (
	"fmt"
	"os"

	"github.com/ew1l/wb-l2/develop/dev10/telnet"
)

func main() {
	t := telnet.NewTelnet()

	if err := t.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}
