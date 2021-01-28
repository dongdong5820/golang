//
// 时间操作大全
//
package main

import (
	"fmt"
	"time"
)

func main() {
	// 1 获取时间
	// 1.1 获取当前时间
	// 返回当前时间，返回的是time.Time类型,UTC时区
	fmt.Println("获取时间===========")
	now := time.Now()
	fmt.Println(now)
	// 当前时间戳(秒级)
	fmt.Println(now.Unix())
	// 纳秒级
	fmt.Println(now.UnixNano())
	// 时间戳小数部分(单位:纳秒)
	fmt.Println(now.Nanosecond())

	// 1.返回当前年月日时分秒、星期几、一年中的第几天
	// 日期
	year, month, day := now.Date()
	fmt.Printf("year:%d, month:%d, day:%d\n", year, month, day)
	// 年
	fmt.Println(now.Year())
	// 月
	fmt.Println(now.Month())
	// 日
	fmt.Println(now.Day())
	// 时分秒
	hour, minute, second := now.Clock()
	fmt.Printf("hour:%d, minute:%d, second:%d\n", hour, minute, second)
	// 星期
	fmt.Println(now.Weekday())
	// 一年中对应的第几天
	fmt.Println(now.YearDay())
	// 时区
	fmt.Println(now.Location())

	// 2.格式化时间
	// "2006-01-02 15:04:05"  2006 1 2 3 4 5
	// 时间戳(int) <-->  time.Time类型 <--> 日期格式(string)
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.Format("2006-01-02"))
	fmt.Println(now.Format("15:04:05"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))

	// 01 时间戳与日期字符相互转化
	layout := "2006-01-02 15:04:05"
	// 0.根据秒，纳秒返回time.Time类型
	t := time.Unix(1610416795, 0)
	fmt.Println(t.Format(layout))
}
