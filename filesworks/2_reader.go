package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer file.Close()

	fileInfo, _ := file.Stat()
	if fileInfo.Size() == 0 {
		fmt.Println("Файл пуст")
		return
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при сканировании файла:", err)
	}
}
