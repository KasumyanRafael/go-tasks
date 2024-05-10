package main

import (
	"fmt"
	"os"
)

func main() {
	fileName := "read-only-file.txt"

	// Создаем файл только для чтения
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDONLY, 0444)
	if err != nil {
		fmt.Println("Ошибка создания файла:", err)
		return
	}
	defer file.Close()

	// Пытаемся записать данные в файл
	_, writeErr := file.WriteString("Эта строка не должна быть записана в файл")
	if writeErr != nil {
		fmt.Println("Попытка записи в файл завершилась ошибкой, как и ожидалось:", writeErr)
	} else {
		fmt.Println("Данные успешно записаны в файл")
	}
}
