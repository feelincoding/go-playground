package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 버퍼 사용하기 위해 선언, os.Stdin은 표준 입력 장치를 의미
	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("입력하세요")
		var number int
		_, err := fmt.Scanln(&number)
		if err != nil {
			fmt.Println(err)
			fmt.Println("숫자를 입력하세요")

			// 키보드 버퍼를 비우는 것
			stdin.ReadString('\n')
			continue
		}
		fmt.Println("입력하신 숫자는", number, "입니다.")
		if number%2 == 0 {
			fmt.Println("짝수입니다.")
			fmt.Println("종료합니다.")
			break
		} else {
			fmt.Println("홀수입니다.")
		}
	}
}
