package main

import "fmt"

func main() {
	var exam1, exam2, exam3, totalScore int

	fmt.Println("Баллы ЕГЭ.")
	fmt.Println("Введите результат первого экзамена:")
	fmt.Scanln(&exam1)

	fmt.Println("Введите результат второго экзамена:")
	fmt.Scanln(&exam2)

	fmt.Println("Введите результат третьего экзамена:")
	fmt.Scanln(&exam3)

	totalScore = exam1 + exam2 + exam3

	fmt.Println("Сумма проходных баллов:", 275)
	fmt.Println("Количество набранных баллов:", totalScore)

	if totalScore >= 275 {
		fmt.Println("Поздравляем, вы поступили!")
	} else {
		fmt.Println("Вы не поступили.")
	}
}
