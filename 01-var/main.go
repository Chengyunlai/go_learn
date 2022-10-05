// 创建包main，声明了一个main包
package main

// 导入内置包文件,fmt可以调用控制台的输出方法
import(
	"fmt"
)
// const PI float64 = iota
const (
	PI = iota
	II
	GI
	)
// main函数程序入口
func main() {
    fmt.Println(II)
}