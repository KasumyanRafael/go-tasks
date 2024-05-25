package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:] // Получаем аргументы командной строки

	if len(args) == 0 {
		fmt.Println("Пожалуйста, укажите имена файлов в качестве аргументов командной строки.")
		return
	}

	if len(args) == 1 {
		fileContent, err := ioutil.ReadFile(args[0])
		if err != nil {
			fmt.Println("Ошибка чтения файла:", err)
			return
		}
		fmt.Println(string(fileContent))
		return
	}

	// Получаем содержимое первого файла
	fileContent1, err := ioutil.ReadFile(args[0])
	if err != nil {
		fmt.Println("Ошибка чтения первого файла:", err)
		return
	}

	// Получаем содержимое второго файла
	fileContent2, err := ioutil.ReadFile(args[1])
	if err != nil {
		fmt.Println("Ошибка чтения второго файла:", err)
		return
	}

	// Конкатенируем содержимое файлов
	concatenatedContent := string(fileContent1) + "\n" + string(fileContent2)

	// Если указан результатный файл, записываем результат туда
	if len(args) == 3 {
		err := ioutil.WriteFile(args[2], []byte(concatenatedContent), 0644)
		if err != nil {
			fmt.Println("Ошибка записи результатного файла:", err)
			return
		}
		fmt.Println("Результат сохранен в", args[2])
		return
	}

	// Выводим результат на экран
	fmt.Println(concatenatedContent)
}

//go run main.go first.txt second.txt result.txt
