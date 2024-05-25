package main

import (
	"fmt"
	"strconv"
)

func main() {
	var ticketNumber int
	fmt.Println("Введите четырехзначный номер билета:")
	fmt.Scanln(&ticketNumber)

	ticketString := strconv.Itoa(ticketNumber)

	if len(ticketString) != 4 {
		fmt.Println("Ошибка! Номер билета должен быть четырехзначным")
		return
	}

	if ticketString[0] == ticketString[3] && ticketString[1] == ticketString[2] {
		fmt.Println("Билет является зеркальным")
	} else {
		sum := int(ticketString[0]-'0') + int(ticketString[1]-'0') - int(ticketString[2]-'0') - int(ticketString[3]-'0')
		if sum == 0 {
			fmt.Println("Билет является счастливым")
		} else {
			fmt.Println("Билет является обычным")
		}
	}
}
