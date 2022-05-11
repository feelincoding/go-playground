package main

import "fmt"

func main() {
	var a int = 10
	b := 3.14
	fmt.Println(float64(a)*b, a*int(b))
}
