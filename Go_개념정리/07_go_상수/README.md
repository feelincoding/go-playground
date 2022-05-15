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
      Red int = iota
      Blue
      Green
  )
  // 선언하는 타입도 int로 같다면 이렇게 생략할 수 있음
  ```
