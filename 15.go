package main

import (
	"strings"
	"wb_1/common"
)

var justString string

func someFunc() {
	v := common.RandomString(1 << 10)
	println(v)

	// Ссылаемся на память, где содержиться большая строка, GC не будет чистить её
	// justString = v[:100]

	// Делаем копию нужной подстроки, при выходе из функции никто не будет на неё ссылаться, поэтому garbage collector
	// отчистит память с ней
	justString = strings.Clone(v[:100])
}

func main() {
	someFunc()
	println(justString)
}
