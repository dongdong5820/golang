// Cf converts its numeric argument to Celsius and Fahrenheit.
// 编译该源代码 go build cf.go
// 执行命令 ./cf.exe 32
package main

import (
	"fmt"
	"os"
	"strconv"
	// 导入自定义包
	"./tempconv"
)

func main()  {
	for _,arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n",
			f, tempconv.FToC(f), c, tempconv.CToF(c))
	}
}

