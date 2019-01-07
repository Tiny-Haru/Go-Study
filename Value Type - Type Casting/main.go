package main

import "fmt"

func main() {
	a := 12345678
	var b = float64(a)
	// 타입을 함수처럼 사용하여 타입 캐스팅이 가능하다.
	_ = b

	fmt.Println(a, b) // 12345678 1.2345678e+07
	fmt.Println(string(65)) // A
}
