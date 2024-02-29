// Go语言的运算符
package main

import "fmt"

func main() {
	//算数运算符
	fmt.Println(10 / 4)    //整数相除，结果有小数的话，小数部分忽略，只保留整数
	fmt.Println(10 % 4)    //取模运算，获取除法的余数
	fmt.Println(-10 % 4)   //取模运算，结果的正负与被模数的符号相同
	fmt.Println(10 % (-4)) //取模运算，结果的正负与被模数的符号相同

	// ++ 和 -- 的使用
	var i int = 10
	i++
	fmt.Println(i) // 输出 11
	i--
	fmt.Println(i) // 输出 10
}
