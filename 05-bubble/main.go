package main

import(
	"fmt"
)

func main() {
	values := []int{4, 3, 14, 85, 34, 27, 91, 95, 26,12,32}
	fmt.Println(values)
	BubblingASC(values)//正序冒泡 从小到大
	BubblingDESC(values)//倒序冒泡 从大到小
}

func BubblingASC(values []int) {
	for i:=0; i<len(values); i++ {
		for j:=i+1; j<len(values); j++ {
			if values[i] > values[j] {
				values[i],values[j] = values[j],values[i]
			}
		}
	}
	fmt.Println(values)
}

func BubblingDESC(values []int) {
	for i:=0; i<len(values); i++ {
		for j:=i+1; j<len(values); j++ {
			if values[i] < values[j] {
				values[i],values[j] = values[j],values[i]
			}
		}
	}
	fmt.Println(values)
}