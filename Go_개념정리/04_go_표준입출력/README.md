# `Go` 표준입출력

- `fmt`라는 패키지로 제공한다.
  - `Print()`: 함수 입력값들을 출력
  - `Println()`: 함수 입력값들을 출력하고 개행
    - 실수 출력시 지수 표현으로 될 때가 있다
    - %v 형태로 알아서 출력
  - `Printf()`: 서식에 맞게 출력
    - 실수 출력시 실수 표현으로 출력된다
- 기타
  - `\n`: 개행
  - `\t`: 탭
  - `\\`: \ 출력
  - `\*`: \*에 해당하는 것을 출력
- 서식지정자(몇가지만)
  - |    구분    |                                      설명                                       |
    | :--------: | :-----------------------------------------------------------------------------: |
    |    `%v`    | 데이터 타입에 맞춰서 기본형태로 출력<br>즉, `%v`쓰면 어지간하면 알아서 출력해줌 |
    |    `%T`    |                               데이터 타입을 출력                                |
    |    `%b`    |                                  이진수로 출력                                  |
    |    `%e`    |                              지수 형태로 실수 출력                              |
    |    `%f`    |                              실수 형태로 실수 출력                              |
    | `%g`, `%G` |                  길이에 따라서 <br>지수형태나 실수 형태로 출력                  |
- %와 서식 사이에 숫자와 부호의 역할

# 표준 입력

- Scan()
- Scanf()
- Scanln(): 제일 많이 쓰임

```go
n, err := fmt.Scanln(&a, &b)
// &: 주소값을 의미
```

# 입력버퍼

- bit stream 형태로 들어간다. 차례차례
- 입력이 실패하면 버퍼를 비워야 된다
- 왜? 실패한 입력 바로 뒷값이 들어가니까 또실패하게 되네? 그러니 비워서 초기화 해야지
