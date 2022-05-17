package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	for {
		if i == 5 {
			break
		}
		fmt.Println(i)
		i++
	}

	for {
		time.Sleep(time.Second)
		fmt.Println(i, time.Second)
		i++
	}
}
