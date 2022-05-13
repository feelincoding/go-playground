package main

import "fmt"

func main() {
	var a int16 = 12345
	var b int8 = int8(a)
	var c int16 = 10
	var d int8 = int8(c)
	fmt.Println(a, b, c, d)
}
