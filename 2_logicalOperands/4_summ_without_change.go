package main

import "fmt"

func main() {
	var coin1, coin2, coin3, price int

	// Запрос пользователю ввести стоимость товара и номиналы трех монет
	fmt.Println("Введите стоимость товара:")
	fmt.Scanln(&price)
	fmt.Println("Введите номиналы трех монет:")
	fmt.Scanln(&coin1)
	fmt.Scanln(&coin2)
	fmt.Scanln(&coin3)

	// Проверка возможности оплаты без сдачи
	if coin1 == price || coin2 == price || coin3 == price ||
		coin1+coin2 == price || coin1+coin3 == price || coin2+coin3 == price ||
		coin1+coin2+coin3 == price {
		fmt.Println("Можно заплатить за товар без сдачи")
	} else {
		fmt.Println("Нельзя заплатить за товар без сдачи")
	}
}
