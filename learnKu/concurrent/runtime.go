//
// runtime包
//
package main

import (
	"fmt"
	"runtime"
	"time"
)

// runtime.Gosched()
func main1() {
	go func(s string) {
		for i := 0; i < 2; i++ {
			fmt.Println(s)
		}
	}("world")
	// 主协程
	for i := 0; i < 2; i++ {
		// 切一下，再次分配任务
		runtime.Gosched()
		fmt.Println("hello")
	}
}

// runtime.Goexit() 退出当前协程
func main() {
	go func() {
		defer fmt.Println("A.defer")
		func() {
			defer fmt.Println("B.defer")
			// 结束协程
			runtime.Goexit()
			defer fmt.Println("C.defer")
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()
	// 主协程
	time.Sleep(time.Second * 3)
}

// Go语言中的操作系统线程和goroutine的关系：
// 1.一个操作系统线程对应用户态多个goroutine
// 2.go程序可以同时使用多个操作系统线程
// 3.goroutine和OS线程是多对多关系，即 m:n
