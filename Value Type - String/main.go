package main

import "fmt"

func main() {
	// 문자열은 "(큰 따옴표) 또는 `(back quoto)를 사용한다.
	//큰 따옴표를 사용한 문자열은 줄바꿈이 불가능하지만, %(이스캐이프 문자열)을 사용할 수 있다.
	fmt.Println("Hello, World") // Hello, World
	fmt.Println(`Hello, World`) // Hello, World

	str := "Hello, World"

	//인덱싱, 덧셈 기호를 통한 결합이 기본적으로 가능하다.
	fmt.Println(str[0]) //72
	fmt.Println(str + "!") //Hello, World!
}
