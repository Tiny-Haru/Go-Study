package main

import "fmt"

func main() {
	var A, B, input int
	fmt.Scanln(&input)
	for i := 0; i < input; i++ {
		fmt.Scanln(&A,&B)
		fmt.Println(A+B)
		A = 0
		B = 0
	}
}
