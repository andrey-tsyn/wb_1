package main

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24}

	println(getIndexOfElement(12, arr))
}

// Бинарный поиск
func getIndexOfElement(val int, arr []int) int {
	start, end := 0, len(arr)-1

	for start <= end {
		midIndex := (start + end) / 2
		midElem := arr[midIndex]

		if midElem == val {
			return midIndex
		} else if midElem < val {
			start = midIndex + 1
		} else {
			end = midIndex - 1
		}
	}

	return -1
}
