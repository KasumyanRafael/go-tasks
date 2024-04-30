package main //задание про улучшение гонщика

import (
	"fmt"
)

func main() {
	var lap int = 3
	var engine int = 256
	var wheels int = 93
	var steeringWheel int = 49
	var wind int = 21
	var rain int = 17

	/* расчет скорости */
	var speed int = engine + wheels + steeringWheel - wind - rain

	fmt.Println("===================")
	fmt.Print("Супергонки. Круг ", lap, "\n")
	fmt.Println("===================")
	fmt.Print("Шумахер (", speed, ")\n")
	fmt.Println("===================")
	fmt.Println("Водитель: Шумахер")
	fmt.Print("Скорость: ", speed, "\n")
	fmt.Println("-------------------")
	fmt.Println("Оснащение")
	fmt.Print("Двигатель: +", engine, "\n")
	fmt.Print("Колёса: +", wheels, "\n")
	fmt.Print("Руль: +", steeringWheel, "\n")
	fmt.Println("-------------------")
	fmt.Println("Действия плохой погоды")
	fmt.Print("Ветер: −", wind, "\n")
	fmt.Print("Дождь: −", rain, "\n")
}
