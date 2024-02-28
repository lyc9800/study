// 指针入门
package main

import "fmt"

func main() {
	var i int = 10
	var ptr *int
	ptr = &i
	fmt.Printf("变量i的地址是: %x\n", ptr)
}
