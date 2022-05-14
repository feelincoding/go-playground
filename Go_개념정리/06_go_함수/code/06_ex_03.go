package main

import "fmt"

func Divide(a int, b int) (result int, success bool) {
	if b == 0 {
		result = 0
		success = false
		return
	} else {
		result = a / b
		success = true
		return
	}
}

func main() {
	c, success := Divide(10, 2)
	fmt.Printf("c = %d, success = %v\n", c, success)
	d, success := Divide(10, 0)
	fmt.Printf("d = %d, success = %v\n", d, success)
}
