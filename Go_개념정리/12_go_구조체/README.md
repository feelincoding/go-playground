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
- embedded field
  - 구조체만 가능하다
  - 근데 이거는 굳이 이렇게 해야하나 싶다, 혼자 개발하는게 아니면 이거는 안쓰는게 가독성 더 좋을 듯?
  - ```go
    type User struct {
      Name string
      Age  int
    }
    type VipUser struct {
      UserInfo User
      VipLevel int
    }
    type VipUser2 struct {
      User // embedded field
      VipLevel int
    }
    ```
- 구조체의 크기
  - 모든 필드의 사이즈 더하기 = 구조체 사이즈
  - Memory Alignment
    - Memory Padding 이 일어나는데 구조체 안에 필드 사이즈중에 가장 큰걸로 통일 되는 듯 왜? 귀찮으니까 그래서 패딩해서 계산하는거지
    - ***`위에거 살짝 잘못 된게 예를 들어 필드 사이즈가 1,8,1,8 이면 패딩에 의해서 8*4 되는데, 1,1,8,8이면 패딩 한번만 일어나서 8*3 된다.`***
- 구조체 복사
  - 모든 필드 값이 복사된다.
  - ~~기본적으로 shallow copy 지원하는 듯, 아닌거 같은데?~~
    - ~~배열은 deep copy가 기본이었는데 신기하네, 2차 배열도 deep copy 일라나?~~
  - 기본적으로 deep copy 지원함! 코드로 확인해 봤다!
- 구조체의 역할
  - **_`low coupling, high cohesion`_**
  - 묶을 것은 묶고(high cohesion, 구조체 역할), 안 묶을 것은 묶지 말고(low coupling)
- ***다른 언어는 구조체와 클래스가 따로지만 `golang 은 클래스가 없다, 즉 구조체가 클래스의 역할을 하며`, 다른 언어는 클래스가 객체로써 역할을 수행한다면 `golang은 구조체가 객체로써의 역할을 수행한다.`***
