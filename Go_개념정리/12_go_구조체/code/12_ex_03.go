package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	a := User{
		Name: "홍길동",
		Age:  20,
	}

	var b = a

	b.Name = "김유신"
	a.Name = "장보고"

	fmt.Println(a)
	fmt.Println(&a.Name, &a.Age)
	fmt.Println(b)
	fmt.Println(&b.Name)
}
