package main

import "fmt"

func printHello() {
	//input : 0, return : 0
	fmt.Println("Hello")
}

func printValue(value int) {
	//input : 1, return : 0
	fmt.Println(value)
}

func returnZero() int {
	//input : 0, return : 1
	return 0
}

func sum(a,b int) int {
	//input : 2, return : 1
	return a+b
}

func main() {
	printHello() //Hello
	printValue(1) //1
	fmt.Println(returnZero())//0
	fmt.Println(sum(1, 2))//3
}