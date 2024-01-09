package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
	"wb_1/common"
)

func main() {
	val, err := common.GetInputFromStdIn("Введите строку")
	if err != nil {
		fmt.Println(err)
		return
	}

	if utf8.RuneCountInString(val) == 0 {
		fmt.Println("Пустая строка")
		return
	}

	if isAllSymbolsUnique(val) {
		fmt.Println("Уникальна")
	} else {
		fmt.Println("Не уникальна")
	}
}

func isAllSymbolsUnique(str string) bool {
	str = strings.ToLower(str)
	m := make(map[rune]struct{})

	for _, sym := range str {
		if _, ok := m[sym]; ok {
			return false
		}

		m[sym] = struct{}{}
	}

	return true
}
