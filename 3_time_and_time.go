package main

import "fmt"

func main() {
	// Входные данные
	duration := 6400
	client_minutes := 350
	cashier_minutes := 700

	// Вычисление полной стоимости товара

	service_time := client_minutes + cashier_minutes

	client_amount := duration / service_time

	// Вывод результата
	fmt.Println("Расчёт количества клиентов, обслуживаемых за смену.")
	fmt.Println("Введите длительность смены в минутах:", duration)
	fmt.Println("Сколько минут клиент делает заказ:", client_minutes)
	fmt.Println("Сколько минут кассир собирает заказ:", cashier_minutes)
	fmt.Println("-----Считаем-----")
	fmt.Println("За смену длиной ", duration, " минут кассир успеет обслужить ", client_amount, " клиентов")
}
