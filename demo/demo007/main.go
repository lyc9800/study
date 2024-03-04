// 进制
// 进制转换
package main

import "fmt"

func main() {
	//十进制转二进制
	var i int = 10
	fmt.Printf("%b\n", i)
	//十进制转八进制
	fmt.Printf("%o\n", i)
	//十进制转十六进制
	fmt.Printf("%x\n", i)
	//位移运算
	//2&3
	fmt.Println(2 & 3)
	//2|3
	fmt.Println(2 | 3)
	//2^3
	fmt.Println(2 ^ 3)
}
