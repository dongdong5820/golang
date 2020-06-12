// Echo3 prints its command-line arguments
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// 内置包 strings.Join 类似php的implode
	fmt.Println(strings.Join(os.Args[1:], " "))
}
