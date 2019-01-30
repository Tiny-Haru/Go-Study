package main

import "fmt"

func main() {
	swap := func(a *int, b *int) {
		*a, *b = *b, *a
	}

	a := new(int)
	b := new(int)
	//new(type)은 타입에 해당하는 메모리를 할당하고, zero value로 초기화한 후 그에 대한 포인터를 반환한다.
	//별도의 초기화 없이 포인터형 변수를 선언(var a *int)하면 nil로 초기화되므로, zero value로 초기화된 포인터 변수를 선언하기 위해 사용된다.
	*a = 1
	*b = 2
	fmt.Println(*a, *b) //1 2
	swap(a, b)
	fmt.Println(*a, *b) //2 1
}
