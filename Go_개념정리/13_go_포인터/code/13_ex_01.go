package main

import "fmt"

func main() {
	a := 10
	fmt.Println("a =", a)
	p := &a
	var p2 *int = &a
	var p3 **int = &p
	fmt.Println("p =", p)
	fmt.Println("p2 =", p2)
	fmt.Println("p3 =", p3)
	**p3 = 20
	fmt.Println("a =", a)
	var p4 *int
	fmt.Println("p4 == nil:", p4 == nil)
}
