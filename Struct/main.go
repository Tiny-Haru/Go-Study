package main

import "fmt"

type position struct {
	//type [name] struct
	//이름이 지정된(name) 필드들로 이루어진 타입
	x int
	y int
	//각 필드들은 이름과 타입을 가진다.
}

func main() {
	var p position
	//구조체는 값 타입이기 때문에 nil로 초기화되지 않고, 필드들이 zero value로 채워진 구조체로 초기화된다.
	//아래와 같이 필드 접근을 할 수 있다.
	fmt.Println(p.x, p.y) //0 0

	p2 := position{x: 2, y: -2}
	fmt.Println(p2) //{2 -2}
	//필드에 대해 값을 할당하며 변수 생성

	p3 := position{3, -3}
	fmt.Println(p3) //{3 -3}
	//필드에 순서대로 할당하는 경우, 필드명을 생략한다.

	p4 := position{} // == var p4 position
	fmt.Println(p4) //{0 0}
	//zero value로 초기화하며 생성

	p5 := new(position) // == var p5 *position, p4 := &position
	fmt.Println(p5) //&{0 0}
	//모든 필드에 메모리가 할당되고, 각 필드들은 zero value로 설정된 후 포인터(*position 타입)가 반환됨.
}
