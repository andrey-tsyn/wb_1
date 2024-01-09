package main

import "fmt"

func main() {
	first := map[string]struct{}{
		"moo":         {}, // Есть во втором
		"meow":        {}, // Есть во втором
		"barking bad": {},
		"quack":       {},
	}

	second := map[string]struct{}{
		"moo":                    {}, // Есть в первом
		"meow":                   {}, // Есть в первом
		"hiss":                   {},
		"what does the fox say?": {},
	}

	intersection := getIntersectionOfSets(first, second)

	fmt.Printf("%v\n", intersection)
}

func getIntersectionOfSets(first, second map[string]struct{}) map[string]struct{} {
	intersection := make(map[string]struct{})

	// Итерируемся по первой мапе и проверяем, есть ли значение с ключем во второй мапе, если да,
	// то добавляем в мапу с пересечениями
	for k := range first {
		if _, ok := second[k]; ok {
			intersection[k] = struct{}{}
		}
	}

	return intersection
}
