package main

import (
	"fmt"
	"time"
)

func main() {
	// 前n年时间
	//str := "2011-01-01 23:59:59"
	//currentTime, _ := time.ParseInLocation("2006-01-02 15:04:05", str, time.Local)
	currentTime := time.Now()
	oldTime := currentTime.AddDate(-4, 0, 0)
	fmt.Printf("%s\n %s\n", currentTime.Format("2006-01-02 15:04:05"), oldTime.Format("2006-01-02 15:04:05"))
}
