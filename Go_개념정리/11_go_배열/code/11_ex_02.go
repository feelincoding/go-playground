package main

import "fmt"

func main() {
	a := [5]int{5, 4, 3, 2, 1}

	b := a
	// var b [5]int = a

	var c [5]int
	for i, v := range a {
		c[i] = v
	}
	fmt.Println(a)
	fmt.Println(&a[0])
	fmt.Println(b)
	fmt.Println(&b[0])
	fmt.Println(c)
	fmt.Println(&c[0])

	fmt.Println("--------------------------------")
	a[0] = 1000
	b[1] = 2000
	fmt.Println(a)
	fmt.Println(&a[0])
	fmt.Println(b)
	fmt.Println(&b[0])
	fmt.Println(c)
	fmt.Println(&c[0])

	d := [2][3]int{{1, 2, 3}, {4, 5, 6}}

	for _, arr := range d {
		for _, v := range arr {
			fmt.Println(v)
		}
	}
}
