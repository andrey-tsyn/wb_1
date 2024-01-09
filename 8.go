package main

import "fmt"

func main() {
	var number int64
	var pos uint
	var bitValue uint

	fmt.Println("Введите целое число: ")
	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Введите позицию бита: ")
	_, err = fmt.Scan(&pos)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Введите значение бита: ")
	_, err = fmt.Scan(&bitValue)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Допустим, что мы хотим изменить число 7 (0111), поменяв значение бита во 2-й позиции
	mask := ^(int64(1) << pos) // Создаем маску (1101 в нашем случае)

	// Не изменяем значения всех битов, кроме нужного, нужный будет иметь значение 0.
	// Это нужно для того, чтобы гарантированно сработала следующая операция.
	number &= mask // В нашем случае получим (0101)

	//Устанавливаем значение бита в заданное на заданной позиции, например 1 (итог 0111) или 0 (итог 0101)
	number = number | (int64(bitValue) << pos)

	fmt.Printf("Результат: %d. Двоичная запись: %b\n", number, number)
}
