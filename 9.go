package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	// Создаем 2 канала по условию
	numsChan := make(chan int, 3)
	multipleChan := make(chan int, 3)

	// Добавляем значения в numsChan из массива
	for _, val := range arr {
		go func(val int, numsChan chan<- int) {
			numsChan <- val
		}(val, numsChan)
	}

	// Получаем значения из numsChan, преобразуем и отправляем в multipleChan
	go func(numsChan <-chan int, multipleChan chan<- int) {
		for val := range numsChan {
			multipleChan <- val * 2
		}
	}(numsChan, multipleChan)

	// Выводим значения, получаемые из multipleChan
	go func(multipleChan <-chan int) {
		for val := range multipleChan {
			fmt.Println(val)
		}
	}(multipleChan)

	<-interrupt
	os.Exit(1)
}
