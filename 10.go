package main

import "fmt"

func main() {
	temperatureMeasures := [...]float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	m := make(map[int][]float32)

	for _, val := range temperatureMeasures {
		// Приводим float32 к int, отбрасывается дробная часть, делим и умножаем на 10, чтобы избавиться 1-го разряда
		group := int(val) / 10 * 10

		// Если в мапе нет значения по ключу, то инициализируем с новым слайсом, иначе просто добавляем к существующему
		// слайсу значение
		s, ok := m[group]
		if !ok {
			s = []float32{val}
			m[group] = s
		} else {
			m[group] = append(s, val)
		}
	}

	for k, v := range m {
		fmt.Printf("%d: %v\n", k, v)
	}
}
