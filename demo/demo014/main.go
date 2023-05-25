package main

import (
	"fmt"
)

func main() {

	fmt.Println("请输入您的成绩:")
	var score int
	fmt.Scanln(&score)

	switch {
	case score >= 90 && score <= 100:
		fmt.Println("您的成绩优秀")
	case score >= 70 && score < 90:
		fmt.Println("您的成绩良好")
	case score > 70 && score <= 60:
		fmt.Println("您的成绩及格")
	case score < 60:
		fmt.Println("您的成绩不及格")
	default:
		fmt.Println("您的成绩无效")
	}
}

