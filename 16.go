package main

import "fmt"

func main() {
	arr := []int{1, 56, 2, 456, 3, 7, 243, 234, 5, 72, 7, 5, 234, 2, 345, 23, 6, 247, 3456, 7, 3, 34, 6, 2, 6754, 67,
		5, 2345, 7, 568, 4678, 6, 3245, 32, 51, 35, 46, 758, 4, 7245, 62456, 568, 678, 4567, 357, 2456, 2436, 87}

	arr = quickSort(arr)

	fmt.Printf("%v\n", arr)
}

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]
	var less, greater []int

	for i := 1; i < len(arr); i++ {
		val := arr[i]

		if val < pivot {
			less = append(less, val)
		} else {
			greater = append(greater, val)
		}
	}

	result := append(quickSort(less), pivot)
	result = append(result, quickSort(greater)...)

	return result
}
