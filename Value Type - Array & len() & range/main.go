package main

import "fmt"

func main() {
	var x [5]int
	y := [5]int{1,2,3,4,5}
	//위와 같이 배열 선언 가능
	//초기화를 하지 않은 배열은 zero value로 초기화
	//x 배열의 경우 int 이므로, {0, 0, 0, 0, 0}으로 초기화됨.

	x[0] = 1
	x[1] = 2

	fmt.Println(len(x))
	//x 배열의 길이

	total := 0

	for _, value := range x {
		total += value
		//range (array) 를 통해 배열을 순회할 수 있다.
		//순회할 때 마다 인덱스와 값을 반환한다.
		//Go는 사용하지 않는 변수가 있으면 컴파일 오류를 발생.
		//언더 스코어(_)를 통해 오류 해결.
	}
	fmt.Println(total)
	fmt.Println(y)
}
