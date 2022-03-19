package main

var justString string

// Проблема заключается в кодировке, символы могут быть unicode. Если символ 'Z' занимает один байт, то символ 'Ш' - два байта
// Решение: конвертировать всю строку в rune
func someFunc() {
	v := createHugeString(1 << 10)

	// Конвертируем строку в rune
	u := []rune(v)

	justString = string(u[:100])

	// fmt.Println(justString)
	// До изменения:
	// ШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШ
	// ZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZ
	// После:
	// ШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШШ
	// ZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZ
}

func createHugeString(size int) string {
	var v string
	for i := 0; i < size; i++ {
		// v += "Z"
		v += "Ш"
	}

	return v
}

func main() {
	someFunc()
}
