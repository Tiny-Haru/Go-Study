package main

import "fmt"

func plma(a, b int) (int, int) {
	//Go는 반환값을 명시하는 부분에 (type, type, ...)형태로 타입을 명시하면 여러개의 값을 반환할 수 있다.

	return a + b, a - b
}

func main() {
	plus, minus := plma(10,5)
	//반환값을 순서대로 변수에 대입

	fmt.Println(plus, minus)
}
