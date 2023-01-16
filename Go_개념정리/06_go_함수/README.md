# `Go` 함수

- 함수의 정의 순서

  - 함수정의키워드 / 함수명 / 매개변수 / 반환타입

  ```go
  func funcName(a int, b int) int {
      return a+b
  }
  ```

- 함수콜은 스택구조!
- `Go` 에서는 멀티 반환이 가능하다! return 여러개, 여러타입
- 여러개 선언을 할때는 그냥 `:=` 을 쓰게 해준다!!
  - ```go
    c, success := Divide(10, 2)
    fmt.Printf("c = %d, success = %v\n", c, success)
    d, success := Divide(10, 0)
    fmt.Printf("d = %d, success = %v\n", d, success)
    // success는 선언이 되었었지만 d가 새롭게 선언되는 녀석으로
    // 걍 := 을 쓰게해준다!
    ```
- `return`할 변수명을 명시하지 않는다면 일일히 `return`해야한다.
  - example
  ```go
  func Divide(a int, b int) (int, bool) {
    if b == 0 {
  	  return 0, false
    } else {
  	  return a / b, true
    }
  }
  ```
- `return`할 변수명을 명시함으로써 편하게 `return`할 수 있다.
  - example
  ```go
  func Divide(a int, b int) (result int, success bool) {
    if b == 0 {
  	  result = 0
  	  success = false
  	  return
    } else {
  	  result = a / b
  	  success = true
  	  return
    }
  }
  ```
- 재귀호출
