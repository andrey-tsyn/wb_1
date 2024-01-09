package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	index := 4

	if index < len(s) && index >= 0 {
		s = append(s[:index], s[index+1:]...)
	}

	fmt.Printf("%v\n", s)
}
