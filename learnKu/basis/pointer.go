package main

import "fmt"

func main() {
	var i1 = 5
	fmt.Printf("An integer:%d, its location in memory: %p\n", i1, &i1)
	// 定义一个整形指针
	var intP *int
	intP = &i1
	fmt.Printf("The value at memory location %p is %d\n", intP, *intP)

	// 字符串指针
	s := "good bye"
	var p *string = &s
	*p = "ciao"
	// prints address
	fmt.Printf("Here is the pointer p: %p\n", p) 
	// prints string
	fmt.Printf("Here is the string *p: %s\n", *p)
	// prints the same string
	fmt.Printf("Here is the string s: %s\n", s)
}

