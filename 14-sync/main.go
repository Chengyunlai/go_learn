package main

import (
	"fmt"
	"sync"
	"time"
)

var x int = 1
var wg sync.WaitGroup
var lock sync.Mutex

func buy(name string) {
	lock.Lock()
	if x > 0 {
		time.Sleep(time.Second)
		x = x - 1
		fmt.Printf("%s，买到了包\n", name)
		lock.Unlock()
	} else {
		lock.Unlock()
	}

	defer func() {
		wg.Done()
	}()
}
func main() {
	wg.Add(3)
	go buy("小王")
	go buy("小明")
	go buy("张三")
	wg.Wait()
}
