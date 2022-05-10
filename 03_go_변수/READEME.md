# `Go`의 변수
```go
package main

import "fmt"

func main() {
	var a int = 10
    // 변수를선언한다 변수명 타입 = 넣을데이터
	var msg string = "Hello Variable"

	a = 20
    // a 라는 메모리 공간에 있는 데이터 수정
	msg = "Good Mornig"
    // msg 라는 메모리 공간에 있는 데이터 수정
	fmt.Println(a, msg)
}
```