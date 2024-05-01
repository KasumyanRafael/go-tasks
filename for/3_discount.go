package main

import (
	"fmt"
)

func main() {
	var price, discount, maxDiscount float64

	fmt.Println("Введите цену товара:")
	fmt.Scanln(&price)

	fmt.Println("Введите скидку в процентах:")
	fmt.Scanln(&discount)

	maxDiscountPercent := 30.0
	maxDiscountAmount := 2000.0

	if discount > maxDiscountPercent {
		discount = maxDiscountPercent
	}

	maxDiscount = price * (discount / 100)
	if maxDiscount > maxDiscountAmount {
		maxDiscount = maxDiscountAmount
	}

	fmt.Println("Сумма скидки:", maxDiscount)
}
