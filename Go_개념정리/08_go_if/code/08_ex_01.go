package main

import (
	"bufio"
	"fmt"
	"os"
)

func checkFunc(a int) (int, bool) {
	if a == 0 {
		return 0, false
	} else {
		return a, true
	}
}

func main() {
	var a int
	var b int

	stdin := bufio.NewReader(os.Stdin)

	if n, error := fmt.Scanln(&a); error == nil && n == 1 {
		if result, check := checkFunc(a); check {
			fmt.Println("올바른 int 입력, a: ", result, "!!")
		} else {
			fmt.Println("0입력했군")
		}
	} else {
		fmt.Println("마지막찬스다")
		stdin.ReadString('\n')
		if n, error = fmt.Scanln(&b); error == nil && n == 1 {
			if result, check := checkFunc(b); check {
				fmt.Println("올바른 int 입력, b: ", result)
			} else {
				fmt.Println("0입력했군")
			}
		} else {
			fmt.Println("넌 탈락")
			stdin.ReadString('\n')
		}
	}
}
