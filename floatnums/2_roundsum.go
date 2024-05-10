package main

import (
	"fmt"
	"math"
)

func main() {
	var deposit float64
	var monthlyInterestRate float64
	var years int

	fmt.Print("Введите сумму вклада: ")
	fmt.Scan(&deposit)

	fmt.Print("Введите ежемесячный процент капитализации: ")
	fmt.Scan(&monthlyInterestRate)

	fmt.Print("Введите количество лет: ")
	fmt.Scan(&years)

	totalInterest := 0.0

	for i := 1; i <= years*12; i++ {
		interest := deposit * monthlyInterestRate / 100
		totalInterest += interest
		deposit += math.Floor(interest * 100)
	}

	totalAmount := deposit + totalInterest
	roundingAmount := totalAmount - math.Floor(totalAmount)

	fmt.Printf("Итоговая сумма на руках вкладчика: %.2f\n", math.Floor(totalAmount))
	fmt.Printf("Сумма, зачисленная в пользу банка за счет округления копеек: %.2f\n", roundingAmount)
}
