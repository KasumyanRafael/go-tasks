package main

import "fmt"

func wrapFunction(fn func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		defer func() {
			fmt.Println("Функция завершила свою работу")
		}()
		return fn(a, b)
	}
}

func main() {
	// Анонимные функции, которые будут обернуты и вызваны с помощью wrapFunction
	add := func(a, b int) int {
		fmt.Println("Сложение выполнено")
		return a + b
	}

	subtract := func(a, b int) int {
		fmt.Println("Вычитание выполнено")
		return a - b
	}

	multiply := func(a, b int) int {
		fmt.Println("Умножение выполнено")
		return a * b
	}

	// Вызовем wrapFunction с каждой из анонимных функций
	wrappedAdd := wrapFunction(add)
	fmt.Println("Результат сложения:", wrappedAdd(5, 3))

	wrappedSubtract := wrapFunction(subtract)
	fmt.Println("Результат вычитания:", wrappedSubtract(10, 4))

	wrappedMultiply := wrapFunction(multiply)
	fmt.Println("Результат умножения:", wrappedMultiply(7, 2))
}
