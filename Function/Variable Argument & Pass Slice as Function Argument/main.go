package main

import "fmt"

func sum(values ...int) int {
	//[name] ...[type]
	//가변인자는 자바스크립트의 rest parameter와 비슷한 형식으로 표현
	total := 0

	for _, value := range values {
		total += value
	}
}

func main() {
	fmt.Println(sum([]int{1, 3, 5, 7, 9}...))
	//[slice]...
	//슬라이스를 인자로 풀어서 전달할 수 있다.
}
