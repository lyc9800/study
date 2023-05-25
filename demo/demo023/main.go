package main

import (
	"fmt"
)

func main() {
	//100以内的数求和，求出第一次和大于20时候跳出
	var sum int
	for count := 1; count <= 100; count++ {
		sum += count
		if sum >= 20 {
			break
		}
		fmt.Println(count)
	}
	//实现登录，有三次机会，判断用户名称为：张无忌，密码为8888，提示登录成功，否则输出还有几次机会
	var name string
	var passWord int32
	var conut int = 3
	for i := 3; i >= 1; i-- {
		fmt.Println("请输入用户名称:")
		fmt.Scanln(&name)

		fmt.Println("请输入密码:")
		fmt.Scanln(&passWord)
		conut--
		if name == "张无忌" && passWord == 8888 {
			fmt.Println("登录成功")
			break
		} else if conut == 0 {
			fmt.Println("您的账号已经被锁定")
		} else {
			fmt.Printf("您还有%v次机会\n", conut)
		}
	}

}
