package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	// Считываем значение с STDIN и преобразуем в int
	fmt.Println("Введите кол-во секунд перед завершением работы: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	input := scanner.Text()

	secs, err := strconv.ParseInt(input, 10, 0)
	if err != nil {
		log.Fatal(err)
	}

	// Создаем таймер на заданное кол-во секунд
	timer := time.NewTimer(time.Duration(secs) * time.Second)

	intCh := make(chan int, 1)

	go func() {
		for {
			time.Sleep(500 * time.Millisecond)
			intCh <- rand.Int()
		}
	}()

	go func() {
		for val := range intCh {
			fmt.Println(val)
		}
	}()

	// Ждём сигнала от таймера, когда закончится время
	<-timer.C
}
