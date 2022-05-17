package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Printf("%d번째 도는중\n", i+1)
	}

	j := 0
	for j < 5 {
		fmt.Print(j, ", ")
		j++
	}
	fmt.Print(j, "\n")
}
