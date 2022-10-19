package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	var in chan []int = make(chan []int, 10)
	var out chan []int = make(chan []int, 10)
	x := make([]int, 0)
	x = append(x, 1, 2, 3, 4, 5, 6, 7)
	out <- x
	close(out) // 关闭通道

	wg.Add(1)
	go odd(in, out)

	for v := range in { //同样需要告诉通道关闭了
		fmt.Println(v)
	}
	//time.Sleep(time.Second * 2000)
	wg.Wait()
}

func odd(in chan<- []int, out <-chan []int) {
	tmp := make([]int, 0)
	for res := range out { // 用range需要告诉通道关闭了
		//fmt.Println(res)
		for _, value := range res {
			if value%2 == 0 {
				//fmt.Println(value)
				tmp = append(tmp, value)
			}
		}
	}
	in <- tmp
	close(in)
	wg.Done()
}

//1. chan<- int是一个只能发送的通道，可以发送但是不能接收；
//2. <-chan int是一个只能接收的通道，可以接收但是不能发送。
