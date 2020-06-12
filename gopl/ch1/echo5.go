// Echo5 prints its command-line arguments
package main

import (
	"fmt"
	"os"
)

func main() {
	for i, s := range os.Args[1:] {
		fmt.Printf("%d\t%s\n", (i + 1), s)
	}
}
