package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go func() {
		fmt.Println("并发程序")
		wg.Done()
	}()
	fmt.Println("主程序")
	wg.Wait()
}
