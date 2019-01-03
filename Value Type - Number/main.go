package main

import "fmt"

func main() {
	//자료형중 숫자는 int와 float로 구분된다.

	a := 1
	b := 1.0

	//정수는 int8, 16, 32, 64, uint8, 16, 32, 64를 사용할 수 있다.
	var c uint64 = 18446744073709551615

	//uint8와 byte가 같고, int32와 rune은 같다.
	var d byte = 255
	var e rune = 2147483647

	//float에는 float32, 64가 있다.
	//부동소수점을 표현하기 위한 complex64, 128이 존재한다.
	var f float32 = 1.129391199
	var g complex64 = 3i + 1.1312

	fmt.Println(a, b, c, d, e, f, g)
}
