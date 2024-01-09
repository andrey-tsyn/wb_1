package main

import "fmt"

func main() {
	a, b := 1, 2
	fmt.Printf("Before swap. a = %d, b = %d\n", a, b)
	a, b = b, a
	fmt.Printf("After swap. a = %d, b = %d\n", a, b)
}
