package main

import (
	"time"
	"fmt"
)

func producer(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
		//[chan] <- [value]
		//채널에 값 적재하기
		time.Sleep(time.Microsecond * 45)
	}
}

func consumer(c chan int) {
	for {
		msg := <- c
		//[variable] = <- [value]
		//채널에서 값을 가져오며, 가져올 값이 없다면 생길 때 까지 wait 한다.

		fmt.Println(msg)
	}
}

func main() {
	c := make(chan int)
	//make(chan [type])
	//채널은 두 고루틴이 서로 통신하고 실행 흐름을 동기화하는 수단을 제공한다.

	go producer(c)
	go consumer(c)

	time.Sleep(time.Second * 5)
}
