package main

import (
	"fmt"
)

func main() {
	str := "Hello 월드"
	for i := 0; i < len(str); i++ {
		fmt.Printf("타입: %T 값: %d 문자값: %c\n", str[i], str[i], str[i])
	}
	// str2 := "안녕하세요"
	// fmt.Printf("%c\n", str2[6])
	fmt.Print("=====")
	fmt.Print(len(str))
	fmt.Print("=====")
	arr := []rune(str)
	fmt.Print("=====")
	fmt.Print(len(arr))
	fmt.Print("=====")
	for i := 0; i < len(arr); i++ {
		fmt.Printf("타입: %T 값: %d 문자값: %c\n", arr[i], arr[i], arr[i])
	}
}
