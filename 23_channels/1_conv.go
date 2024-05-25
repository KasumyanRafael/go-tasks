package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
)

func squareWorker(input <-chan int, output chan<- int) {
	for num := range input {
		square := num * num
		output <- square
	}
	close(output)
}

func productWorker(input <-chan int, output chan<- int, wg *sync.WaitGroup) {
	for square := range input {
		product := square * 2
		output <- product
	}
	wg.Done()
}

func main() {
	input := make(chan int)
	squares := make(chan int)
	products := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)

	go squareWorker(input, squares)
	go productWorker(squares, products, &wg)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {
			var val string
			fmt.Scan(&val)
			if val == "стоп" {
				close(input)
				break
			}
			num, err := strconv.Atoi(val)
			if err == nil {
				input <- num
			} else {
				fmt.Println("Пожалуйста, введите число или 'стоп' для завершения.")
			}
		}
	}()

	go func() {
		for product := range products {
			fmt.Println("Произведение:", product)
		}
	}()

	<-signals // Ждем сигнала остановки

	wg.Wait()
	close(squares)
	close(products)

	fmt.Println("Программа завершена.")
}
