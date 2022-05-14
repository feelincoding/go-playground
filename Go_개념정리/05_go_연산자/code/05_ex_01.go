package main

import "fmt"

func main() {
	var a int8 = 2
	var b int8 = 64
	fmt.Printf("a = %d, a(2) = %08b, a<<2 = %d, a<<2(2) = %08b\n", a, a, a<<2, a<<2)
	fmt.Printf("b = %d, b(2) = %08b, b<<2 = %d, b<<2(2) = %08b\n", b, b, b<<2, b<<2)
}
