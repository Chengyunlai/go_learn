package main

import "fmt"

func main() {
	num := 0
	notUsePoint(num)
	fmt.Println(num) //0
	usePoint(&num)
	fmt.Println(num) //10
}

func usePoint(num *int) {
	*num = 10
}

func notUsePoint(num int) {
	num = 10
}
