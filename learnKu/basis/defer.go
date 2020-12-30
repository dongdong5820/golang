package main

import "fmt"

func main() {
	InvokeDeferFunction()
}

func InvokeDeferFunction() {
	defer DeferFunctionCall()
	defer OtherDeferFunctionCall()
	fmt.Println("Still executing InvokeDeferFunction")
}

func DeferFunctionCall() {
	fmt.Println("Defer Function called")
}

func OtherDeferFunctionCall() {
	fmt.Println("Other Defer Function called")
}
