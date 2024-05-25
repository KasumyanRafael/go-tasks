package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	go func() {
		for i := 1; ; i++ {
			select {
			case <-interrupt:
				fmt.Println("Получен сигнал завершения, дождитесь завершения текущей задачи...")
				time.Sleep(3 * time.Second)
				fmt.Println("Выход из программы")
				os.Exit(0)
			default:
				square := i * i
				fmt.Println("Квадрат числа", i, "равен", square)
				time.Sleep(1 * time.Second)
			}
		}
	}()

	fmt.Println("Нажмите Ctrl+C для завершения работы")
	select {}
}
