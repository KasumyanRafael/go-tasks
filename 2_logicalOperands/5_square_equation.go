package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64

	// Запрос пользователю ввести коэффициенты квадратного уравнения
	fmt.Println("Введите коэффициент a:")
	fmt.Scanln(&a)
	fmt.Println("Введите коэффициент b:")
	fmt.Scanln(&b)
	fmt.Println("Введите коэффициент c:")
	fmt.Scanln(&c)

	// Вычисление дискриминанта
	discriminant := b*b - 4*a*c

	// Проверка значения дискриминанта для нахождения корней уравнения
	if discriminant > 0 {
		root1 := (-b + math.Sqrt(discriminant)) / (2 * a)
		root2 := (-b - math.Sqrt(discriminant)) / (2 * a)
		fmt.Println("У уравнения два различных корня:")
		fmt.Println("Корень 1:", root1)
		fmt.Println("Корень 2:", root2)
	} else if discriminant == 0 {
		root := -b / (2 * a)
		fmt.Println("У уравнения один корень:", root)
	} else {
		fmt.Println("У уравнения нет действительных корней")
	}
}
