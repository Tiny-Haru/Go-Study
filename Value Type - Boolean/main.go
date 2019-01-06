package main

import "fmt"

func main() {
	//True와 False를 나타내는 타입인 Boolean에 대해 알아보자.
	a := true
	b := false

	fmt.Println(a && b) //a와 b를 AND 연산 한 결과
	fmt.Println(a || b) //a와 b를 OR 연산 한 결과
	fmt.Println(!a, !b) //a와 b에 각각 NOT 연산을 한 결과
}
