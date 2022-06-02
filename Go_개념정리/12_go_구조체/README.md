# `Go` 구조체

- 구조체의 선언

  - 여러 필드를 묶어서 하나의 구조체로 만들 수 있다
  - 예제

    ```go
    type Student struct {
      // type 타입명 struct
      Name string
      // 필드명 타입
      Class int
      No int
      Score float64
    }

    var a Student

    // interface
    type ... interface{}
    ```

- 구조체의 변수 초기화
  - 처음 선언을 하면 모든 필드 값들은 기본값으로 초기화 된다.
  - 초기화시 특정 필드만 초기화 할 수도 있다.
  - `golang의 특징` 초기화 할 때 한 줄이 아닌 여러줄로 초기화 할 때는 마지막 값 뒤에 쉼표를 넣어야 한다.
    - ```go
      var house House = House{ "서울시", 28, 10, "아파트" }
      var house House = House{
         "서울시",
         28,
         10,
         "아파트",
      }
      ```
- 구조체를 포함하는 구조체 ex) 학교안에 학생과 교사/ 교사, 학생이라는 구조체가 모여서 학교를 이룬다.
