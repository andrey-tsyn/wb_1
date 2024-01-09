package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"wb_1/common"
)

func main() {
	// Прослушиваем сигнал на завершение работы программы
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	// Считываем значение с STDIN и преобразуем в int
	fmt.Println("Введите кол-во воркеров: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	input := scanner.Text()

	workerCount, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}

	// Создаем буферизированный канал и wait group
	randStrsChanel := make(chan string, 10)
	wg := sync.WaitGroup{}

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func(ch <-chan string) {
			defer wg.Done()

			for {
				val, ok := <-ch
				// Если канал закрыт и не осталось значений, то завершаем выполнение
				if !ok {
					fmt.Println("Воркер завершил работу, канал закрыт.")
					break
				}

				fmt.Println(val)
			}
		}(randStrsChanel)
	}

Main:
	for {
		select {
		case <-interrupt:
			// Закрываем канал и ожидаем завершения работы всех горутин. При закрытом канале все горутины считают переменную,
			// которая показывает, закрыт ли он.
			// Выбрал этот метод, т.к. он прост в реализации, позволяет уведомить всех о завешреннии работы и
			// корректно завершить каждую горутину.
			close(randStrsChanel)
			wg.Wait()
			break Main
		default:
			// Создаем рандомную строку и отправляем в канал
			randStrsChanel <- common.RandomString(10)
		}
	}

	// Other logic...
}
