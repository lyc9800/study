package main

import "fmt"

func main() {
	//按照输入的a,b,c,d,e,f  来显示对应的周一到周日
	//switch语句完成
	var key byte

	fmt.Println("请输入一个字符: a,b,c,d,e,f,g")
	fmt.Scanf("%c", &key)

	switch key {
	case 'a':
		fmt.Println("您选择的是周一")
	case 'b':
		fmt.Println("您选择的是周二")
	case 'c':
		fmt.Println("您选择的是周三")
	case 'd':
		fmt.Println("您选择的是周四")
	case 'e':
		fmt.Println("您选择的是周五")
	case 'f':
		fmt.Println("您选择的是周六")
	case 'g':
		fmt.Println("您选择的是周日")
	default:
		fmt.Println("输入有误，程序退出")
	}
}
