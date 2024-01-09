package main

import (
	"fmt"
	"sync"
	"wb_1/common"
)

func main() {
	nums := [...]int{2, 4, 6, 8, 10}
	sum := 0
	// Создаем буферизированный канал
	squaredNumsCh := make(chan int, 3)

	go func() {
		wg := sync.WaitGroup{}

		for _, num := range nums {
			wg.Add(1)
			go func(num int, ch chan<- int) {
				defer wg.Done()

				// Возводим значение в квадрат и отправляем в канал
				val := common.Square(num)
				ch <- val
			}(num, squaredNumsCh)
		}

		// Ожидаем выполнение всех горутин, после закрываем канал, если закрыть канал и попробовать записать
		// после - будет паника
		wg.Wait()
		close(squaredNumsCh)
	}()

	// Цикл читает входящие из канала значения, если канал закрыт и не осталось значений в буфере, то цикл завершается
	for val := range squaredNumsCh {
		sum += val
	}

	fmt.Printf("Sum is: %d\n", sum)
}
