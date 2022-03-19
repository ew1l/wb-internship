package main

import "fmt"

func main() {
	temps := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// Map, где будет хранится результат
	temps_map := make(map[int][]float32)

	for _, temp := range temps {
		// Ключ должен быть округлен до ближайшего десятка. Например, -43 => -40 и тд
		key := int(temp) / 10 * 10
		temps_map[key] = append(temps_map[key], temp)
	}

	fmt.Println(temps_map)
}

// map[-20:[-25.4 -27 -21] 10:[13 19 15.5] 20:[24.5] 30:[32.5]]
