// 读取键盘输入
package main

import "fmt"

func main() {
	//要求：从控制台读取用户信息
	//方式1：fmt.Scanln
	var name string
	var age int
	var address string
	var email string
	fmt.Println("请输入姓名:")
	fmt.Scanln(&name)
	fmt.Println("请输入年龄:")
	fmt.Scanln(&age)
	fmt.Println("请输入地址:")
	fmt.Scanln(&address)
	fmt.Println("请输入邮箱:")
	fmt.Scanln(&email)
	fmt.Printf("姓名：%v,年龄：%v,地址：%v,邮箱：%v\n", name, age, address, email)
}
