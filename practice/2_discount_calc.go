package main

import "fmt"

func main() {
	const max_discount = 10000
	const percent = 0.1

	var weekday int
	var sum float64

	fmt.Println("Введите номер дня недели (от 1 до 7):")
	fmt.Scanln(&weekday)

	fmt.Println("Введите сумму вашего чека:")
	fmt.Scanln(&sum)

	if weekday >= 1 && weekday <= 5 { // Проверка на будний день
		discount := 0.0
		if sum >= max_discount {
			discount = sum * percent
			fmt.Printf("Вам предоставляется скидка в размере %.2f руб.\n", discount) //округление до сотых
		} else {
			fmt.Println("Сумма вашего чека не достигла порога для получения скидки.")
		}
	} else {
		fmt.Println("В данное время скидка не действует.")
	}
}
