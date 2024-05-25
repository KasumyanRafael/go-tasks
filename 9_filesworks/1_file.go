package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	file, err := os.Create("log.txt")
	if err != nil {
		fmt.Println("Не удалось создать файл:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Введите строку (или 'exit' для выхода): ")
		scanner.Scan()
		line := scanner.Text()

		if line == "exit" {
			fmt.Println("Программа завершена.")
			break
		}

		now := time.Now()
		timeStr := now.Format("2006-01-02 15:04:05")
		data := fmt.Sprintf("%s %s\n", timeStr, line)

		_, err = writer.WriteString(data)
		if err != nil {
			fmt.Println("Ошибка записи в файл:", err)
			break
		}
	}
}
