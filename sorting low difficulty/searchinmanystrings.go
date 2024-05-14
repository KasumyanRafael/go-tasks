package main

import (
	"fmt"
	"strings"
)

func parseTest(sentences []string, chars []rune) [][]int {
	result := make([][]int, len(sentences))

	for i, sent := range sentences {
		words := strings.Fields(sent)
		lastWord := words[len(words)-1]

		result[i] = make([]int, len(chars))

		for j, char := range chars {
			index := strings.LastIndex(string(lastWord), string(char))
			result[i][j] = index
		}
	}

	return result
}

func main() {
	sentences := []string{"Пример предложения с разными символами: 123", "Ещё одно предложение 456"}
	chars := []rune{'а', 'о', 'е', 'и', 'у', '1', '2', '3'}

	output := parseTest(sentences, chars)

	for i, row := range output {
		fmt.Printf("Предложение %d:\n", i+1)
		for j, index := range row {
			fmt.Printf("Индекс символа %c: %d\n", chars[j], index)
		}
		fmt.Println()
	}
}
