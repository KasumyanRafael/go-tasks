package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	text := "a10 10 20b 20 30c30 30 dd"
	var numStr string

	for _, char := range text {
		if unicode.IsDigit(char) || char == ' ' {
			numStr += string(char)
		} else {
			if numStr != "" {
				num, err := strconv.Atoi(numStr)
				if err == nil {
					fmt.Println("Часть строки, которая может быть приведена к числу:", num)
				}
				numStr = ""
			}
		}
	}

	// Check the last part of the string
	if numStr != "" {
		num, err := strconv.Atoi(numStr)
		if err == nil {
			fmt.Println("Часть строки, которая может быть приведена к числу:", num)
		}
	}
}
