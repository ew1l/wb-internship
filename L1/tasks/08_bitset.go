package main

import "fmt"

func main() {
	var value int64 = 12 // == 1100 [1(3) 1(2) 0(1) 0(0)]
	i := 0

	value |= 1 << i
	// 12 == 1100
	// 1 << 0 == 1 * 2^0 = 1
	// 12 | 1 = 1100 | 0001 = 1101 = 13

	fmt.Println(value)
}
