package main
import (
	"fmt"
	"runtime"
)

var prompt = "Enter a digit, e.g. 3 " + "or %s to quit."

func init() {
	if runtime.GOOS == "windows" {
		prompt = fmt.Sprintf(prompt, "Ctrl+Z, Enter")
	} else {
		prompt = fmt.Sprintf(prompt, "Ctrl+D")
	}
}

func main() {
	bool1 := true
	if bool1 {
		fmt.Printf("The value is true\n")
	} else {
		fmt.Printf("The value is false\n")
	}
	fmt.Printf(prompt + "\n")
	fmt.Printf("%d\n", Abc(-5))
}

func Abc(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func isGreater(x, y int) bool {
	if x > y {
		return true
	}
	return false
}