package main

func main() {
	stops := []string{"Проспект Доватора", "Алые Паруса", "улица Московская", "улица Весенняя 2"}
	pass_in := []int{7, 4, 8, 16}
	pass_out := []int{0, 3, 1, 8}
	ticketcost := 20
	totalprofit := 0
	pass_in_tram := 0
	wages := 80
	fuel := 64
	taxes := 64
	repairment := 64
	for i := 0; i < 4; i++ {
		println("Прибываем на остановку ", stops[i])
		println("В салоне пассажиров: ", pass_in_tram)
		println("Сколько пассажиров зашло на остановке: ", pass_in[i])
		println("Сколько пассажиров вышло: ", pass_out[i])
		println("Отправляемся с остановки  ", stops[i])
		pass_in_tram = pass_in_tram + pass_in[i]
		pass_in_tram = pass_in_tram - pass_out[i]
		totalprofit = totalprofit + (pass_in_tram * ticketcost)
	}
	println("Всего заработано ", totalprofit, " руб.")
	println("Зарплата водителя: ", wages, " руб.")
	println("Расходы на топливо: ", fuel, " руб.")
	println("Налоги: ", taxes, " руб.")
	println("Ремонт машин: ", repairment, " руб.")
	totalprofit = totalprofit - fuel - taxes - wages - repairment
	println("Итого доход: ", totalprofit, " руб.")

}
