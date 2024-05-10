package main

import (
	"fmt"
	"strings"
)

func main() {
	days := map[string]string{
		"пн": "понедельник",
		"вт": "вторник",
		"ср": "среда",
		"чт": "четверг",
		"пт": "пятница",
	}

	var input string
	fmt.Print("Введите сокращенное название буднего дня (пн, вт, ср, чт, пт): ")
	fmt.Scanln(&input)

	day, ok := days[strings.ToLower(input)]
	if !ok {
		fmt.Println("Некорректный ввод.")
		return
	}

	fmt.Println("Следующие рабочие дни:")
	for i := 0; i < 5; i++ {
		fmt.Println(day)
		if day == "пятница" {
			break
		}
		for key, value := range days {
			if value == day {
				day = value
				break
			}
		}
	}
}
