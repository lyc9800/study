// 练习
package main

import "fmt"

func main() {
	//假如还有97天放假，请问是几个星期几天
	days := 97
	weeks := days / 7
	remainder := days % 7
	fmt.Printf("还有%d天,也就是%d个星期%d天\n", days, weeks, remainder)

	//逻辑运算符
	var age int = 18
	if age > 20 && age < 30 {
		fmt.Println("young")
	} else if age >= 10 && age <= 20 {
		fmt.Println("ok1")
	}

	if age > 20 || age < 10 {
		fmt.Println("ok2")
	} else if age > 10 || age < 100 {
		fmt.Println("ok3")
	}

	//赋值运算符
	var num1 int = 10
	var num2 int = 20
	var num3 int
	num3 = num1
	num1 = num2
	num2 = num3
	fmt.Printf("num1=%d ,num2=%d ,num3=%d", num1, num2, num3)
}
