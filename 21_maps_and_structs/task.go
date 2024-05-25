package main

import (
	"fmt"
)

// Структура Student
type Student struct {
	name  string
	age   int
	grade int
}

func main() {
	studentMap := make(map[string]*Student) // Инициализируем хранилище map[studentName]*Student

	for {
		var name string
		var age, grade int
		fmt.Println("Введите имя, возраст и оценку студента:")
		_, err := fmt.Scanf("%s %d %d\n", &name, &age, &grade) // Считываем данные с stdin

		if err != nil {
			break // Выйти из цикла при получении EOF (ctrl + d)
		}

		student := newStudent(name, age, grade) // Создаем нового студента
		studentMap[name] = student              // Сохраняем указатель на студента в хранилище
	}

	fmt.Println("Студенты из хранилища:")
	for name, student := range studentMap {
		fmt.Printf("%s %d %d\n", student.name, student.age, student.grade) // Выводим информацию о студентах
	}
}

// Функция для создания новой структуры Student
func newStudent(name string, age int, grade int) *Student {
	return &Student{name, age, grade}
}
