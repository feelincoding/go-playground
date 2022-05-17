package main

import "fmt"

func main() {
	// day := "Mon"
	day := "Wed"
	switch day {
	case "Mon", "Tue":
		fmt.Println("월 화 중에 골랐네")
	case "Thu", "Fri":
		fmt.Println("목 금 중에 골랐네")
	default:
		fmt.Println("나머지")
	}

	temp := 18
	switch true {
	case temp < 10, temp > 30:
		fmt.Println("걍 집에 있으세요")
	case temp > 15 && temp < 25:
		fmt.Println("것 옷 걸치면 좋은 날씨")
	case temp > 0 && temp <= 18:
		fmt.Println("애매해")
	}

	// switch dayday := "Fri"; dayday { // 이렇게 if처럼 초기화랑 동시에 써줄 수 있다.
	switch dayday := "Mon"; dayday { // 이렇게 if처럼 초기화랑 동시에 써줄 수 있다.
	case "Mon", "Tue":
		fmt.Println("월 화 중에 골랐네")
	case "Thu", "Fri":
		fmt.Println("목 금 중에 골랐네")
		fallthrough
	default:
		fmt.Println("나머지")
	}
}
