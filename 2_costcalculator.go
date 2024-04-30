package main

import "fmt"

func main() {
	// Входные данные
	product_cost := 6400
	deliver_cost := 350
	discount := 700

	// Вычисление полной стоимости товара

	real_cost := product_cost + deliver_cost - discount

	// Вывод результата
	fmt.Println("Калькулятор стоимости товара со скидкой.")
	fmt.Println("Стоимость товара:", product_cost)
	fmt.Println("Стоимость доставки:", deliver_cost)
	fmt.Println("Размер скидки:", discount)
	fmt.Println("---------")
	fmt.Println("Итого:", real_cost)
}
