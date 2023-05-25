package main

import "fmt"

func main() {
	var height int
	var wealth float32 =1.0
	var look bool

	fmt.Println("请输入您的身高(cm):")
	fmt.Scanln(&height)

	fmt.Println("请输入您的财富(千万):")
	fmt.Scanln(&wealth)

	fmt.Println("请输入您的样貌(帅--true;丑--flase):")
	fmt.Scanln(&look)

	if height >= 180 && wealth >= 1.0 && look  {
		fmt.Println("我一定嫁给你……")
	} else if height == 180 || wealth >= 1.0 || look {
		fmt.Println("嫁吧，比上不足比下有余……")
	} else {
		fmt.Println("滚……")
	}
}
