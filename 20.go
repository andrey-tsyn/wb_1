package main

import "strings"

func main() {
	str := "snow dog sun"

	strs := strings.Split(str, " ")

	for i, j := 0, len(strs)-1; i < j; i, j = i+1, j-1 {
		strs[i], strs[j] = strs[j], strs[i]
	}

	resultStr := strings.Join(strs, " ")

	println(resultStr)
}
