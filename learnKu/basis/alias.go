package main

import (
	"fmt"
	fm "fmt"
)

func main() {
	// 定义常量方式
	const beef, two, c = "ead", 2, "veg"
	// cosnt Monday, Tuesday, Wednesday, Thursday, Friday, Saturday = 1, 2, 3, 4, 5, 6
	const (
		Monday, Tuesday, Wednesday = 1, 2, 3
		Thursday, Friday, Saturday = 4, 5, 6
	)
	const (
		Unknown = 0
		Female  = 1
		man     = 2
	)
	fm.Println("hello, world, alias")

	// 类型转换
	var n int16 = 34
	var m int32
	// compiler error:  cannot use n (type int16) as type int32 in assignment
	// m = n
	m = int32(n)
	fmt.Println("32 bit int is: %d\n", m)
	fmt.Println("16 bit int is: %d\n", n)
}
