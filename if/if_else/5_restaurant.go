package main

import "fmt"

func main() {
	var dayOfWeek, numOfGuests, totalBill int
	var discount, serviceCharge int

	fmt.Println("Введите день недели (1 - понедельник, 2 - вторник и т.д.):")
	fmt.Scanln(&dayOfWeek)

	fmt.Println("Введите число гостей:")
	fmt.Scanln(&numOfGuests)

	fmt.Println("Введите сумму чека:")
	fmt.Scanln(&totalBill)

	discount = 0
	serviceCharge = 0

	if dayOfWeek == 1 {
		discount = totalBill * 10 / 100
	}

	if dayOfWeek == 5 && totalBill > 10000 {
		discount += totalBill * 5 / 100
	}

	if numOfGuests > 5 {
		serviceCharge = totalBill * 10 / 100
	}

	netTotal := totalBill - discount + serviceCharge

	fmt.Println("Скидка по пятницам:", discount)
	fmt.Println("Надбавка на обслуживание:", serviceCharge)
	fmt.Println("Сумма к оплате:", netTotal)
}
