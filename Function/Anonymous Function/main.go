package main

import "fmt"

func main() {
	//Go는 다른 언어들의 많은 장점들을 수용하고자 했다.
	//따라서 Javascript의 대표적인 특징인 익명 함수를 지원한다.

	func(a, b int, callback func(int)) {
		callback(a + b)
	}(1, 3, func(result int) {
		fmt.Println(result) //4
	})

	//두 ㅍ수를 더해 콜백에 전달하는 함수를 익명으로 생성 후 즉시 실행.
}
