//
// 正则
//
package main

import (
	"fmt"
	"log"
	"regexp"
)

func main() {

	data := "<a href='www.baidu.com'>I am a good man</a>"
	fmt.Println(data)

	// 正则替换
	re, err := regexp.Compile("<[A|a][^>]*?>|</[A|a]>")
	if err != nil {
		log.Fatal(err)
	}
	rep := re.ReplaceAllString(data, "")
	fmt.Println(rep)
}
