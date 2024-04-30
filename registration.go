package main

import "fmt"

func main() {
	var username string
	var password string
	var age int

	username = "me"

	password = "Aa123456"

	age = 20

	fmt.Printf("Поздравляю, %s, теперь вы зарегистрированы!\n", username)
	fmt.Printf("Ваш пароль: %s\n", password)
	fmt.Printf("Ваш возраст: %d\n", age)
}
