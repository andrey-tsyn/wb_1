package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	// Первый способ: закрытие канала
	ch := make(chan struct{})

	go func() {
		for {
			select {
			case _, ok := <-ch:
				if !ok {
					return
				}

			// Больше case'ов для, например, работы с другими каналами

			default:
				fmt.Println("Do something...")
			}
			time.Sleep(10 * time.Millisecond)
		}

		// Или:
		/*
			for val := range ch {
				// Some logic
			}
		*/
	}()

	time.Sleep(100 * time.Millisecond)
	close(ch)

	// Второй способ: использование контекста

	// context.WithCancel возращает контекст и функцию для отмены
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			// Если функция отмены вызвана, канал done закрывается, в case обрабатываем
			// случай получения данных или закрытия канала.
			case <-ctx.Done():
				return
			default:
				fmt.Println("Do something...")
			}

			time.Sleep(100 * time.Millisecond)
		}
	}(ctx)

	go func() {
		time.Sleep(3 * time.Second)
		// Вызываем функцию отмены
		cancel()
	}()

	fmt.Println("Main end")

	<-interrupt
}
