// Echo1 prints its command-line arguments
// go build echo1.go  得到 echo1.exe
// 执行命令 ./echo1.exe a b c d e f ==> output a b c d e f
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
