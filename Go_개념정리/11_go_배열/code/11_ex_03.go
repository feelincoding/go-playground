package main

import "fmt"

func main() {
	var arr [5]int = [5]int{5, 4, 3, 2, 1}
	var nums [5]int
	days := [3]string{"Mon", "Tue", "Wed"}
	var temps [5]float64 = [5]float64{12.1, 15.3}
	var s = [5]int{1: 10, 3: 30}
	x := [...]int{10, 20, 30}

	fmt.Println(arr);
	fmt.Println(nums);
	fmt.Println(days);
	fmt.Println(temps);
	fmt.Println(s);
	fmt.Println(x);
}
