package main

import "fmt"

func main() {
	a:=0
	//C 또는 JAVA와 비슷한 구조를 가지고 있음.
	if a==1 {
		fmt.Println("a == 1")
	} else if a == 0 {
		fmt.Println("a == 0")
	} else {
		fmt.Println("else")
	}

	//for의 초기화처럼, 조건문 앞에 Statement를 실행할 수 있음.
	if b := 123; b == 123 {
		fmt.Println("b ==", b)
	}
}
