package main //складывание в уме

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	num1 := rand.Intn(100)
	num2 := rand.Intn(100)

	fmt.Printf("Сколько будет %d плюс %d? Введите ответ: ", num1, num2)
	var answer int
	fmt.Scanln(&answer)

	correctAnswer := num1 + num2

	if answer == correctAnswer {
		fmt.Println("Правильно! Вы молодец!")
	} else {
		fmt.Printf("Неправильно! Правильный ответ: %d\n", correctAnswer)
	}
}
