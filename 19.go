package main

func main() {
	str := "главрыба"
	strRunes := []rune(str)

	for i, j := 0, len(strRunes)-1; i < j; i, j = i+1, j-1 {
		strRunes[i], strRunes[j] = strRunes[j], strRunes[i]
	}

	resultStr := string(strRunes)

	println(resultStr)
}
