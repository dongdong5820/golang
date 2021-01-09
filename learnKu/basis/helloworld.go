package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("hello world")
	fmt.Println("Go Go Go !!!")
	str := "hello, 你好"
	fmt.Println(utf8.RuneCountInString(str))
}
