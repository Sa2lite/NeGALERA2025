package main

import (
	"fmt"
	"strings"
)

func main() {
	sentence := "это молитва, используемая в Римско-католической церкви, в которой просят Бога об освобождении душ верующих из Чистилища."

	words := strings.Fields(sentence)

	for i, word := range words {
		fmt.Printf("%d: %s\n", i+1, word)
	}
}
