# `Go` switch

- 기본 스위치문 형태: 타 언어랑 비슷한듯?

  ```go
  switch a {
      case 1:
       sentence
      case 2:
       sentence
      default:
       문장
  }
  // 이런형태로 진행 a 자리에는 비교할 수 있는 값이 들어가는거니까 스트링 이런 것도 되겠다
  ```
- `case에서 하나 걸리면 다른 언어는 보통 break 해줘야하는데 Go는 케이스 하나 실행하면 알아서 스위치 빠져나온다!`
- `그러면 다른 케이스도 이어서 실행하고 싶으면 어떻게해? -> fallthrough 써주면 됨 그러면 다음 케이스까지 실행됨`, ***`근데 이게 다음게 참인지 거짓인지 범위에 속하는지 안속하는지 관계없이 실행되니까 쓰는거 지양하자`***
- if 문을 좀 더 보기 좋게 써서 좋은 것 같긴한데 굳이라는 생각이 든다, 실제로 tdd에서도 switch 지양 하라는 말도 있긴 했었고
- 하나의 case에서 하나 이상의 값을 `,` 를 통해서 비교할 수 있다.
  ```go
  day := "Mon"
  switch day { // day:="Mon"; day 이렇게 해도 됨, 초기화 식에 함수 써도 ㄱㅊ
  case "Mon", "Tue":
  	fmt.Println("월 화 중에 골랐네")
  case "Thu", "Fri":
  	fmt.Println("목 금 중에 골랐네")
  default:
  	fmt.Println("나머지")
  }
  // 이 때 케이스들끼리 중 겹치는게 있으면 안된다.
  ```
- 변수 넣어서 값만 비교하는게 아니라 true, false가 되는 조건문을 검사도 가능하다.
  ```go
  temp:=18
  switch true { // true 안 쓰면 default가 true임
  case temp<10, temp>30:
  	fmt.Println("걍 집에 있으세요")
  case temp>15 && temp<25:
  	fmt.Println("것 옷 걸치면 좋은 날씨")
  case temp>0 && temp<=18:// 위 케이스에서 걸리니까 여기부터는 실행 안되지
  	fmt.Println("애매해")
  }
  ```
- const 열거값을 통해서도 switch 작성 가능, 09_ex_02.go 참조