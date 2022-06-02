package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	Age   int32
	Score float64
}
type User2 struct {
	Age   int32
	Score float32
}
type User3 struct {
	Test  int8
	Age   int32
	Score float64
}

func main() {
	user := User{29, 100.0}
	user2 := User2{29, 100.0}
	user3 := User3{1, 29, 100.0}

	fmt.Println(unsafe.Sizeof(user.Age))
	fmt.Println(unsafe.Sizeof(user.Score))
	fmt.Println(unsafe.Sizeof(user))
	fmt.Println(unsafe.Sizeof(user2.Age))
	fmt.Println(unsafe.Sizeof(user2.Score))
	fmt.Println(unsafe.Sizeof(user2))
	fmt.Println(unsafe.Sizeof(user3.Test))
	fmt.Println(unsafe.Sizeof(user3.Age))
	fmt.Println(unsafe.Sizeof(user3.Score))
	fmt.Println(unsafe.Sizeof(user3))
}
