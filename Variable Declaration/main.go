package main

import "fmt"

func main() {
	var a int
	//자료형 선언

	var b int = 1
	//선언과 동시에 초기화

	var c = 2
	//타입 추론

	var d, e int
	//여러 변수 선언

	f:=5
	//':='를 통해 var와 타입을 한 번에 명시 가능.

	fmt.Println(a, b, c, d, e, f)
	//초기화 하지 않은 변수는 0으로 출력
}