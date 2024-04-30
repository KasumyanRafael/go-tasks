package main

import (
	"fmt"
)

func main() {
	var height int
	var growth int
	var consumespeed int
	var days int

	fmt.Print("Высота саженца?")
	fmt.Scanln(&height)
	fmt.Print("Скорость роста?")
	fmt.Scanln(&growth)
	fmt.Print("Скорость поедания?")
	fmt.Scanln(&consumespeed)
	fmt.Print("Количество дней?")
	fmt.Scanln(&days)
	println("К середине ", days, " дня бамбук вырастет до ", height+growth*days-consumespeed*days, " сантиметров")

}
