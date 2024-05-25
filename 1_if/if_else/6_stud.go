package main

import "fmt"

func main() {
	var N, K, index int

	// Вводим количество студентов N и количество групп K
	fmt.Print("Введите количество студентов N: ")
	fmt.Scanln(&N)

	fmt.Print("Введите количество групп K: ")
	fmt.Scanln(&K)

	// Вводим порядковый номер студента
	fmt.Print("Введите порядковый номер студента: ")
	fmt.Scanln(&index)

	// Определяем номер группы для студента
	group := (index-1)%K + 1

	fmt.Println("Студент с порядковым номером", index, "попадает в группу номер", group)
}
