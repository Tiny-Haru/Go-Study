package main

import "fmt"

func main() {
	//초기화 없이 선언된 변수들이 기본값으로 채워지는데, 이것을 zero value 라고 부름
	var a int
	var b float32
	var c complex64
	var d string
	var e bool
	var s []int
	var m map[int] string

	fmt.Println("=====Let's study Zero Value=====")
	fmt.Println("int = " , a) //0
	fmt.Println("float32 = ", b) //0
	fmt.Println("complex64 = ", c) //(0+0i)
	fmt.Println("string = ", d) //
	fmt.Println("bool = ", e) //false
	fmt.Println("s == nil :", s==nil, ", slice = ", s) //nil
	fmt.Println("m == map :", m==nil, ", map =", m)//nil
	//nil은 초기화되지 않은 slice, map, 포인터, 인터페이스 등에서 사용되는 zero value 이다.
}
