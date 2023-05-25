package main

import "fmt"

func main() {
	var name string
	var age int
	var address string
	var isPass bool

	fmt.Println("请输入名字:")
	fmt.Scanln(&name)

	fmt.Println("请输入年龄")
	fmt.Scanln(&age)

	fmt.Println("请输入地址:")
	fmt.Scanln(&address)

	fmt.Println("是否通过考试:")
	fmt.Scanln(&isPass)

	fmt.Printf("您的名字是: %v\n您的年龄是: %v\n您的地址是: %v\n您是否通过考试: %v", name, age, address, isPass)
}
