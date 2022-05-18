package main

import "fmt"

func useFlag() {
	a := 2
	b := 1
	found := false
	for ; a < 10; a++ {
		for b = 1; b < 10; b++ {
			fmt.Printf("%d * %d = %d\n", a, b, a*b)
			if a*b == 45 {
				found = true
				break
			}
		}
		if found {
			break
		}
	}
}
func useLable() {
	a := 2
	b := 1
OuterLoop:
	for ; a < 10; a++ {
		for b = 1; b < 10; b++ {
			fmt.Printf("%d * %d = %d\n", a, b, a*b)
			if a*b == 45 {
				break OuterLoop
			}
		}
	}
}
func useCleanCodeFunc() {
	a := 2
	for ; a < 10; a++ {
		var found bool
		if found = find45(a); found {
			break
		}
	}
}
func find45(a int) bool {
	for b := 1; b < 10; b++ {
		fmt.Printf("%d * %d = %d\n", a, b, a*b)
		if a*b == 45 {
			return true
		}
	}
	return false
}
func main() {
	// Flag를 이용하여 중첩 for 문 탈출
	// useFlag()
	// Flag를 이용하여 중첩 for 문 탈출
	// useLable()
	// Flag를 이용하여 중첩 for 문 탈출
	useCleanCodeFunc()
}
