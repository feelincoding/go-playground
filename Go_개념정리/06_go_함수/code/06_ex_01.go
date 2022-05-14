package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

func main() {
	var a int = 1
	var b int = 2
	var c int = Add(a, b)

	fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)
}
