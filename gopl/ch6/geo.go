package main

import (
	"./geometry"
	"fmt"
)

func main() {
	p := geometry.Point{1,2}
	q := geometry.Point{4,6}
	// 调用geometry包中的函数
	fmt.Println(geometry.Distance(p,q))
	// 调用geometry包中Point类型的方法
	fmt.Println(p.Distance(q))
	// 调用新方法，计算三角形的周长
	perim := geometry.Path{
		{1,1},
		{5,1},
		{5,4},
		{1,1},
	}
	fmt.Println(perim.Distance()) // "12"
}
