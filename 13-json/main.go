package main

import (
	"encoding/json"
	"fmt"
)

// 结构体
type Prescription struct {
	Name     string `json:"name"`
	Unit     string
	Additive *Prescription `json:"additive,omitempty"`
}

func main() {
	p := Prescription{}
	p.Name = "鹤顶红"
	p.Unit = "1.2kg"
	p.Additive = &Prescription{
		Name: "砒霜",
		Unit: "0.5kg",
	}

	buf, err := json.Marshal(p) //转换为json返回两个结果
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	fmt.Println("json = ", string(buf))

	// json 字符串转为结构体
	jsonstr := `{"name":"鹤顶红","unit":"1.2kg","additive":{"name":"砒霜","unit":"0.5kg"}}`
	//var res Prescription
	res := Prescription{}
	fmt.Println(&res)
	if err := json.Unmarshal([]byte(jsonstr), &res); err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
