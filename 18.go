package main

import (
	"fmt"
	"sync"
)

func main() {
	counter := Counter{
		mu:    sync.Mutex{},
		value: 0,
	}

	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.AddValue(1)
		}()
	}

	wg.Wait()

	fmt.Printf("%d\n", counter.Value())
}

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) AddValue(val int) {
	c.mu.Lock()
	c.value += val
	c.mu.Unlock()
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}
