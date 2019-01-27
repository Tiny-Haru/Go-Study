package main

import "fmt"

func main() {
	first := func() {
		fmt.Println("First")
	}

	second := func() {
		fmt.Println("Second")
	}
	//Go언어에는 defer라는 특별한 statement가 존재한다.
	defer second()
	first()

	//defer는 특정 문장, 혹은 함수를 나중에 실행할 수 있게 해준다.
	//따라서 위의 실행 순서는 first() -> second()이다.
}
