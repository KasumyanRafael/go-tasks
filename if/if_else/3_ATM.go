package main

import "fmt"

func main() {
	var withdrawalAmount int

	fmt.Println("Банкомат.")
	fmt.Println("Введите сумму снятия со счёта:")
	fmt.Scanln(&withdrawalAmount)

	if withdrawalAmount > 100000 {
		fmt.Println("Превышен лимит суммы для снятия.")
	} else if withdrawalAmount%100 != 0 {
		fmt.Println("Сумма должна быть кратна 100.")
	} else {
		fmt.Println("Операция успешно выполнена.")
		fmt.Println("Вы сняли", withdrawalAmount, "рублей.")
	}
}
