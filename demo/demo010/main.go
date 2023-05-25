package main

import "fmt"

func main() {
	var age int
	fmt.Println("请输入你的年龄:")
	fmt.Scanln(&age)
	if age > 18 {
		fmt.Println("你已经成年了……")
	} else {
		fmt.Println("你还没有成年……请好好学习")
	}
}
