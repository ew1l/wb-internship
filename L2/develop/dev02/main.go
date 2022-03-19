package main

import (
	"fmt"

	"github.com/ew1l/wb-l2/develop/dev02/unpack"
)

func main() {
	set := []string{"a4bc2d5e", "abcd", "45", `qwe\4\5`}
	for _, s := range set {
		us, err := unpack.Unpack(s)
		if err != nil {
			fmt.Printf("%s => %s\n", s, err.Error())
			continue
		}

		fmt.Printf("%s => %s\n", s, us)
	}
}

// a4bc2d5e => aaaabccddddde
// abcd => abcd
// 45 => некорректная строка
// qwe\4\5 => qwe45
