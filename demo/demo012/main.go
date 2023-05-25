package main

import (
	"fmt"
)

func main() {
	//出票系统：4-10月：
	//	旺季
	//	成人（18-60）：60
	//	儿童（<18）:半价
	//	老人（>60）：1/3
	//	淡季：
	//	成人：40
	//	其他：15

	var month int
	var age int
	var price float32 = 60.0
	fmt.Println("请输入您要出行的月份:")
	fmt.Scanln(&month)

	fmt.Println("请输入您的年龄:")
	fmt.Scanln(&age)

	if month >= 4 && month <= 10 {
		if age > 60 {
			fmt.Printf("您的门票是: %v", price/3)
		} else if age < 18 {
			fmt.Printf("您的门票是: %v", price/2)
		} else {
			fmt.Printf("您的票价是: %v", price)
		}
	} else {
		if age >= 18 && age < 60 {
			fmt.Println("您的门票是40")
		} else {
			fmt.Println("您的门票是15")
		}
	}
}
