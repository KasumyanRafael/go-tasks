package main

import "fmt"

func main() {
	// Входные данные
	totaltax := 4000000
	entrances := 10
	flatsInEntrance := 40

	// Вычисление полной стоимости товара

	taxForFlat := totaltax / (flatsInEntrance * entrances)

	// Вывод результата
	fmt.Println("Расчёт стоимости ремонта для каждой квартиры.")
	fmt.Println("Сумма, указанная в квитанции: ", totaltax, " рублей")
	fmt.Println("Подъездов в доме: ", entrances)
	fmt.Println("Квартир в каждом подъезде: ", flatsInEntrance)
	fmt.Println("----Результат-----")
	fmt.Println("Каждая квартира должна платить по ", taxForFlat, " рублей.")
}
