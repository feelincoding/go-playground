# `Go` 상수

- 상수: 변하지 않는 수
- 그럼 어떻게 선언해? `const` 로!
- `iota` 이게 뭐야?: 변수마다 const 일일히 선언 안하고 열거하면서 하는거
  ```go
  const(
      Red int = iota
      Blue int = iota
      Green int = iota
  )
  // 이렇게 iota라고 해놓으면 알아서 0부터 1씩 증가됨, Red=0, ..., Green=2
  const(
      Red int = iota + 1
      Blue
      Green
  )
  // 선언하는 타입도 int로 같다면 이렇게 생략할 수 있음, 이러면 iota + 1 이니까 1부터 3까지
  ```
- Bitflag
- 타입 없는 상수

  - 타입을 정하지 않으면 상수가 쓰일 때 타입이 결정된다.

    ```go
    const PI = 3.14
    const FloatPI float64 = 3.14

    func main() {
        var a int = PI * 100
        // 이 때 PI*100이 a에 들어갈 때 타입이 결정되므로 오류 ㄴㄴ
        var b int = FloatPI * 100
        // 이미 타입을 float64로 설정했기에 타입 오류
        fmt.Println(a)
        fmt.Println(b)
    }
    ```
