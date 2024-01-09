package main

import (
	"errors"
	"fmt"
	"math/big"
	"wb_1/common"
)

func main() {
	// Получаем и парсим первое число
	aStr, err := common.GetInputFromStdIn("Первое число: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	a := new(big.Int)

	a, ok := a.SetString(aStr, 10)
	if !ok {
		fmt.Println("Первое число не удалось считать.")
		return
	}

	// Получаем и парсим второе число
	bStr, err := common.GetInputFromStdIn("Второе число: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	b := new(big.Int)
	b, ok = b.SetString(bStr, 10)
	if !ok {
		fmt.Println("Второе число не удалось считать.")
		return
	}

	// Получаем символ соответствующей операции
	operation, err := common.GetInputFromStdIn("Операции: + - * /")
	if err != nil {
		fmt.Println(err)
		return
	}

	operationFunc, err := getOperationFunc(operation)
	if err != nil {
		fmt.Println(err)
		return
	}

	c := operationFunc(a, b)

	fmt.Printf("Результат: %s\n", c.String())
}

// getOperationFunc возращает функцию для выполнения над двумя целыми числами на основе полученной строки
func getOperationFunc(operation string) (func(a, b *big.Int) *big.Int, error) {
	switch operation {
	case "+":
		return plus, nil
	case "-":
		return minus, nil
	case "*":
		return multiple, nil
	case "/":
		return divide, nil
	default:
		return nil, errors.New(fmt.Sprintf("\"%s\" is not supported operator!\n", operation))
	}
}

func plus(a, b *big.Int) *big.Int {
	a.Add(a, b)
	return a
}

func minus(a, b *big.Int) *big.Int {
	a.Sub(a, b)
	return a
}

func multiple(a, b *big.Int) *big.Int {
	a.Mul(a, b)
	return a
}

func divide(a, b *big.Int) *big.Int {
	a.Div(a, b)
	return a
}
