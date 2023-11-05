package main

import "fmt"

func main() {
	fmt.Println(foo())
	fmt.Println(bar())
}

func foo() string {
	return "Hello"
}

func bar() int {
	return 2
}
