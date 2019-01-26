package main

import "fmt"

func numberGenerator() func() int {
	u := 0
	return func() int {
		u++
		return u
	}
	//Go언어도 정적 스코핑 언어이기 때문에 *클로저가 형성될 수 있다.
	//*클로저 : 함수 바깥에 있는 변수를 참조하는 함수값
}

func main() {
	gen := numberGenerator()

	for i := 0; i < 3; i++ {
		fmt.Println(gen())
	}
}
