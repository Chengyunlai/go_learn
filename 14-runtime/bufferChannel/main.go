package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("操作缓冲通道")
	buffInt := make(chan int, 1)
	wg.Add(1)
	go func() {
		y := <-buffInt
		fmt.Println(y)
		defer wg.Done()
	}()
	buffInt <- 10
	wg.Wait()
	//y := <-buffInt
	//fmt.Println(y)
}
