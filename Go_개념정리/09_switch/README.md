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

- if 문을 좀 더 보기 좋게 써서 좋은 것 같긴한데 굳이라는 생각이 든다, 실제로 tdd에서도 switch 지양 하라는 말도 있긴 했었고
