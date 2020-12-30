//
// goroutine
//
package main

import (
	"fmt"
	"sync"
	"time"
)

func hello() {
	fmt.Println("Hello Goroutine")
}

// 无协程
func main1() {
	hello()
	fmt.Println("main goroutine done!")
}

// 启动单个goroutine
func main2() {
	go hello() // 启动另外一个goroutine去执行hello函数
	fmt.Println("main goroutine done!")
	time.Sleep(time.Second)
}

// 启动多个goroutine
var wg sync.WaitGroup

func helloMulti(i int) {
	defer wg.Done()
	fmt.Println("Hello Goroutine!", i)
}
func main3() {
	for i := 0; i < 10; i++ {
		wg.Add(1) // 启动一个goroutine就登记+1
		go helloMulti(i)
	}
	wg.Wait() // 等待所有登记的goroutine都结束

	fmt.Println("main goroutine done!")
}

// 主协程退出，其他任务不再执行
func main() {
	// 合起来写
	go func() {
		i := 0
		for {
			i++
			fmt.Printf("new goroutine: i = %d\n", i)
			time.Sleep(time.Second)
		}
	}()

	i := 0
	for {
		i++
		fmt.Printf("main goroutine: i =%d\n", i)
		time.Sleep(time.Second)
		if i == 2 {
			break // 退出循环
		}
	}
}
