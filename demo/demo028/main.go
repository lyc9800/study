package main

import "fmt"

//斐波那契数列练习
//使用递归的方式调用

func fbn(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fbn(n-1) + fbn(n-2)
	}
}
func main() {
	fmt.Println(fbn(1))
	fmt.Println(fbn(2))
	fmt.Println(fbn(3))
	fmt.Println(fbn(4))
	fmt.Println(fbn(5))
	fmt.Println(fbn(6))
}
