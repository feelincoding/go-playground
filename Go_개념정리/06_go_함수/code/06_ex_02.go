package main

import "fmt"

func Divide(a int, b int) (int, bool) {
	if b == 0 {
		return 0, false
	} else {
		return a / b, true
	}
}

func main() {
	c, success := Divide(10, 2)
	fmt.Printf("c = %d, success = %v\n", c, success)
	d, success := Divide(10, 0)
	fmt.Printf("d = %d, success = %v\n", d, success)
}
