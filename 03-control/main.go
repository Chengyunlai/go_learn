package main

import(
	"fmt"
)

func main(){
	// if and else
	num := 10
	if num > 0{
		fmt.Println("该数大于0")
	}else{
		fmt.Println("该数小于0")
	}

	if num > 10{
		fmt.Println("该数大于10")
	}else if num > 9{
		fmt.Println("该数大于9")
	}
	// if and func

	if res := number(); res!=0 {
		fmt.Println("判断结果：该数不为0")
	}else{
		fmt.Println("判断结果：该数为0")
	}



	// switc:每个case语句中自带break语句
	gateWay := "001"

	switch gateWay{
  	case "001":
			fmt.Println("网关001")
			fmt.Println("网关001 - copy")
		case "002":
			fmt.Println("网关002")
      fmt.Println("网关002 - copy")
		default:
			fmt.Println("无此网关")
	}

	age := 18
	
	switch age > 20{
		case true:
			fmt.Println("年龄大于18")
		case false:
			fmt.Println("年龄大于17")
	}

	coin := "壹元"
	switch coin{
  	case "一元","壹元":
			fmt.Println("都是一元")
		default:
			fmt.Println("不是一元")
	}

	fmt.Println("---")

	//switch的fallthrough:执行后面没有执行的case代码
	switch gateWay{
	case "001":
		fmt.Println("网关001")
		fmt.Println("网关001 - copy")
		fallthrough
	case "002":
		fmt.Println("网关002")
		fmt.Println("网关002 - copy")
	default:
		fmt.Println("无此网关")
	}

	luckClientId := "0001" //假设我们获取一个用户id是0001

	switch luckClientId{
	case "0001": // 这里我们手动为0001编号id设置了丰盛的奖品
		fmt.Println("字节背包")
		fallthrough
	case "0002":
		fmt.Println("榨汁机")
	default:
		fmt.Println("安慰奖-掘金贴纸")
	}

	fmt.Println("---")

	// for - 01
	for i := 0; i < 10; i++ {
    fmt.Println(i)
  }

	// for - 02
	for j:=0; j<10; {
		fmt.Println(j)
		j++
	}

	// for - 03
	k := 0
	for k < 10 {
    fmt.Println(k)
		k++
	}

	// break and continue 不赘述

	// 标签
	falg:
	for i:=0; i<3; i++{
		for j:=0;j<3;j++{
			fmt.Println(j) // 默认是输出三遍 0 1 2
			// break // 加上这个呢是每次输出0就跳出当前循环，如果说我们加上标签呢？
			break falg // 跳出整个循环,goto道理一样，不建议使用，了解即可
		}
	}

}

func number() int{
	return 0
}

