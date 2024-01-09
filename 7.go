package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	storage := MyMap{
		sync.RWMutex{},
		make(map[int]int),
	}

	// Конкурентно заполняем map
	go func() {
		for i := 0; i < 500; i++ {
			storage.Store(i, rand.Int())
		}
	}()
	go func() {
		for i := 500; i < 1000; i++ {
			storage.Store(i, rand.Int())
		}
	}()

	for i := 0; i < 1000; i += 10 {
		fmt.Println(storage.Get(i))
	}
}

type MyMap struct {
	mu   sync.RWMutex
	dict map[int]int
}

func (m *MyMap) Store(key, value int) {
	// Даем возможность записывать значение только одному пользователю
	m.mu.Lock()
	m.dict[key] = value
	m.mu.Unlock()
}

func (m *MyMap) Get(key int) int {
	// Даем возможность считывать значение многим, гарантируя то, что map останется неизменной во время считывания значения
	m.mu.RLock()
	val := m.dict[key]
	m.mu.RUnlock()
	return val
}
