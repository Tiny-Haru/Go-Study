package main

import "fmt"

func main() {
	sum := 0

	for i := 0; i <= 100; i++ {
		sum += i
	}

	fmt.Println(sum) // 0+1+2+...+100의 결과, 5050
	i := 1

	for i < 100 {
		sum += i
		break
	}

	fmt.Println(sum) //sum + 1 = 5051
}
