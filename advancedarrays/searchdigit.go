package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var size int
	fmt.Print("Введите размер массива: ")
	fmt.Scanln(&size)

	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(100) // генерируем случайное число от 0 до 99 и записываем его в массив
	}

	fmt.Println("Сгенерированный массив:", arr)

	var num int
	fmt.Print("Введите число для поиска: ")
	fmt.Scanln(&num)

	found := false
	count := 0
	for _, val := range arr {
		if found {
			count++
		}
		if val == num {
			found = true
		}
	}

	if found {
		fmt.Printf("Чисел в массиве после %d: %d\n", num, count)
	} else {
		fmt.Println("Введенное число не найдено в массиве.")
	}
}
