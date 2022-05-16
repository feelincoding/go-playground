package main

import (
	"fmt"
	"reflect"
)

const PI = 3.14
const FloatPI float64 = 3.14

func main() {
	var a int = PI * 100
	// 이 때 PI의 타입이 결정되므로 오류 ㄴㄴ
	// var b int = FloatPI * 100
	// 이미 타입을 float64로 설정했기에 타입 오류
	var b int = int(FloatPI * 100)
	var c float64 = FloatPI * 100
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.TypeOf(c))
}
