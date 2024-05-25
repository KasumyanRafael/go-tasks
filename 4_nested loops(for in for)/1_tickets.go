package main

import "fmt"

func isPalindrome(value int) bool {
	reversed := 0
	original := value
	for value > 0 {
		remainder := value % 10
		reversed = reversed*10 + remainder
		value /= 10
	}
	return original == reversed
}

func main() {
	fmt.Println("Зеркальные билеты от 100000 до 999999:")
	for ticket := 100000; ticket <= 999999; ticket++ {
		if isPalindrome(ticket) {
			fmt.Println(ticket)
		}
	}
}
