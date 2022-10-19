package main

import (
	"fmt"
	"time"
)

func main() {
	chan1 := make(chan int, 2)

	//发送
	chan1 <- 1
	chan1 <- 2
	//chan1 <- 3
	time.Sleep(time.Second)

	x := <-chan1
	fmt.Println(x)
}
