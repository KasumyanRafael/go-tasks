package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	text := "Go is an Open source programming Language that makes it Easy to build simple, reliable, and efficient Software."

	count := 0
	words := strings.Fields(text)

	for _, word := range words {
		if len(word) > 0 && unicode.IsUpper(rune(word[0])) {
			count++
		}
	}

	fmt.Printf("Количество слов, начинающихся с большой буквы: %d\n", count)
}
