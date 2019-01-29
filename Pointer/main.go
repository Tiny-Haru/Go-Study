package main

import "fmt"

func main() {
	swap := func(a *int, b *int) {
		*a, *b = *b, *a
	}

	a := 5
	b := 3
	aPtr := &a
	bPtr := &b
	fmt.Println("a :", a, "b :", b)
	swap(aPtr, bPtr)
	fmt.Println("a :", a, "b :", b)
	//Go언어는 포인터를 지원한다.
	//포인터는 값이 저장된 메모리상의 위치를 가르킨다.
	//C언어와 동일하게 사용이 가능하다.
	//변수의 주소를 구히가 위해 &를, 포인터 타입 표현과 포인터 역참조를 위해 *를 사용한다.
}
