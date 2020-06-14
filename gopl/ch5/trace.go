package main

import (
	"fmt"
	"log"
	"time"
)

func main()  {
	bigSlowOperation()
	//_ = double(4)
	fmt.Println(triple(4))
}
// defer函数使用示例
func bigSlowOperation() {
	defer trace("bigSlowOperation")()
	// do not forget the extra parentheses
	// ... lots of work...
	time.Sleep(10 * time.Second)
	// simulate slow operating by sleeping
}
// 函数跟踪
func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}
// 利用defer打印出函数的参数和返回值
func double(x int) (result int) {
	defer func() { fmt.Printf("double(%d) = %d\n", x, result)}()
	return x+x
}
func triple(x int) (result int) {
	defer func() {result += x}()
	return double(x)
}