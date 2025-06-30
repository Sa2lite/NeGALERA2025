package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "В других языках это также называется словарем, хеш-картой или ассоциативным массивом. Вот как создать карту:"

	words := strings.Fields(strings.ToLower(text))

	frequency := make(map[string]int)

	for _, word := range words {
		word = strings.Trim(word, ".,!")
		frequency[word]++
	}

	fmt.Println("Частота слов:")
	for word, count := range frequency {
		fmt.Printf("%s: %d\n", word, count)
	}
}
