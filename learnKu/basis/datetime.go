package main

import (
	"fmt"
	"log"
	"regexp"
	"runtime"
	"time"
)

func main() {
	// 前n年时间
	//str := "2011-01-01 23:59:59"
	//currentTime, _ := time.ParseInLocation("2006-01-02 15:04:05", str, time.Local)
	currentTime := time.Now()
	oldTime := currentTime.AddDate(-4, 0, 0)
	fmt.Printf("%s\n %s\n", currentTime.Format("2006-01-02 15:04:05"), oldTime.Format("2006-01-02 15:04:05"))

	ExecuteTimeTest()
}

//
// 函数耗时统计
//
func TimeTrack(start time.Time) {
	elapsed := time.Since(start)

	pc, _, _, _ := runtime.Caller(1)
	fmt.Println(pc)

	funcObj := runtime.FuncForPC(pc)

	runtimeFunc := regexp.MustCompile(`^.*\.(.*)$`)
	name := runtimeFunc.ReplaceAllString(funcObj.Name(), "$1")
	fmt.Println(name)

	log.Println(fmt.Sprintf("%s took %s", name, elapsed))
}

//
// 测试执行时间
//
func ExecuteTimeTest() {
	defer TimeTrack(time.Now())
	time.Sleep(time.Second * 2)
}
