package main

import (
	"fmt"
	"sync"
	"wb_1/common"
)

func main() {
	nums := [...]int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}

	for _, n := range nums {
		// Добавляем 1 к counter'у, тем самым сообщая, что есть не выполненная горутина
		wg.Add(1)
		// Запускаем горутины и выводим квадраты чисел в stdout
		go func(n int) {
			// Уведомляем о завершении горутины (counter -= 1)
			defer wg.Done()

			square := common.Square(n)
			fmt.Println(square)
		}(n)
	}

	// Ждем, пока горутины не выполнятся(выполнены, если значение counter'a == 0)
	wg.Wait()
}
