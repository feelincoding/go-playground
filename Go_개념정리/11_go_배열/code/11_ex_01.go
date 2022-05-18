package main

import "fmt"

func main() {
	var a [5]int = [5]int{5, 4, 3, 2, 1}

	for num := 0; num < len(a); num++ {
		fmt.Println(a[num])
	}
	for i, v := range a {
		fmt.Println(i, v)
	}
	for _, v := range a {
		fmt.Println(v)
	}

	fmt.Println(a)
}
