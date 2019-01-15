package main

import "fmt"

func main() {
	num := 0

	switch num {
	//일반적인 switch문과 동일하게 사용 가능
	//break는 case에 기본 내장
	//fallthrough를 통해 다음 case 실행 가능.
	case 0:
		fmt.Print("0")
		fallthrough
	case 1:
		fmt.Print("1")
	case 2:
		fmt.Print("2")
	case 3:
		fmt.Print("3")
	default:
		fmt.Print("Not 0,1,2,3")
	}

	//0 1
}