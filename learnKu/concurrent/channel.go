//
// channel通道
//
package main

import (
	"fmt"
	"time"
)

func main1() {
	ch := make(chan int)
	ch <- 10 // 无缓冲区通道，阻塞，会报panic错误
	fmt.Println("发送成功")
}

func main2() {
	ch := make(chan int)
	go recv(ch) // 启用goroutine从通道接收值
	ch <- 10
	fmt.Println("发送成功")
}

// 接收数据
func recv(c chan int) {
	ret := <-c
	fmt.Println("接收成功", ret)
}

func main3() {
	ch := make(chan int, 1) // 创建一个容量为1的有缓冲区通道
	ch <- 10
	fmt.Println("发送成功")
}

// close关闭通道
func main4() {
	c := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c) // 关闭通道
	}()

	for {
		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			break
		}
	}
	fmt.Println("main结束")
}

// 如何优雅的从通道循环取值
func main5() {
	start := time.Now()
	ch1 := make(chan int)
	ch2 := make(chan int)
	// 开启goroutine将0-100的数发送到ch1中
	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		// 关闭通道ch1
		close(ch1)
	}()
	// 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
	go func() {
		for {
			// 通道关闭后再取值ok=false
			i, ok := <-ch1
			if !ok {
				break
			}
			ch2 <- i * i
		}
		// 关闭通道ch2
		close(ch2)
	}()
	// 在主goroutine中从ch2中接收值打印
	for i := range ch2 {
		// 通道关闭后会退出for range循环
		fmt.Println(i)
	}

	// 计算程序耗时
	end := time.Now()
	diff := end.Sub(start)
	fmt.Println("start: ", start)
	fmt.Println("end: ", end)
	fmt.Println("consumer: ", diff)
}

// 单向通道(站在程序角度)
// 1. chan <- int 是一个只能发送的通道，可以发送但是不能接收;
// 2. <- chan int 是一个只能接收的通道，可以接收但是不能发送。
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go counter(ch1)
	go squarer(ch2, ch1)
	printer(ch2)
}

// 将0-99的数依次放入通道
func counter(out chan<- int) {
	for i := 0; i < 100; i++ {
		out <- i
	}
	close(out)
}

// 从in通道中取数据，然后计算平方，放入out通道
func squarer(out chan<- int, in <-chan int) {
	for i := range in {
		out <- i * i
	}
	close(out)
}

// 从in通道中取数据，然后打印
func printer(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}
