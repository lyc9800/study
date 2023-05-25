package main

import "fmt"

//定义一个函数，用来交换两个数值的位置

func getPath(n1 *int, n2 *int) {
	t := *n1
	*n1 = *n2
	*n2 = t
}
func main() {
	a := 10
	b := 20
	getPath(&a, &b)
	fmt.Printf("a=%d\nb=%d\n", a, b)
}
