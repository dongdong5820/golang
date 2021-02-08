package main

import (
	"flag"
	"fmt"
	"os"
)

var name string

// 方式1
func init1() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}
func main1() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
		flag.PrintDefaults()
	}
	flag.Parse()
	fmt.Printf("Hello, %s!\n", name)
}

// 方式2
func init2() {
	flag.CommandLine = flag.NewFlagSet("", flag.ExitOnError)
	flag.CommandLine.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
		flag.PrintDefaults()
	}
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}
func main2() {
	flag.Parse()
	fmt.Printf("Hello, %s!\n", name)
}

// 方式3
var cmdLine = flag.NewFlagSet("question", flag.ExitOnError)

func init() {
	cmdLine.StringVar(&name, "name", "everyone", "The greeting object.")
}
func main() {
	cmdLine.Parse(os.Args[1:])
	fmt.Printf("Hello, %s!\n", name)
}
