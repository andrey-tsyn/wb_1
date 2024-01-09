package main

import (
	"fmt"
	"reflect"
)

func main() {
	integer := 1
	boolean := true
	str := "sfsf"
	ch := make(chan int)

	fmt.Printf("%s\n", getNameOfType(integer))
	fmt.Printf("%s\n", getNameOfType(boolean))
	fmt.Printf("%s\n", getNameOfType(str))
	fmt.Printf("%s\n", getNameOfType(struct{}{}))
	fmt.Printf("%s\n", getNameOfType(ch))
}

func getNameOfType(obj interface{}) string {
	switch obj.(type) {
	case int:
		return "Integer"
	case bool:
		return "Boolean"
	case string:
		return "String"
	// Привязываюсь к конкретному типу, но работает незначительно быстрее
	//case chan int:
	//	return "Channel"

	// Не работает
	//case chan interface{}:
	//	return "Channel"

	default:
		// Не привязываемся к конкретному типу у канала
		t := reflect.TypeOf(obj)

		if t.Kind() == reflect.Chan {
			return "Channel"
		}

		return "Unknown"
	}
}
