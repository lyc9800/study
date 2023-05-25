package main

import "fmt"

//已知：f(1)=3;f(n)=2*f(n-1)+1
//编写程序，求出f(n)的值?
func f(n int) int {
	if n == 1 {
		return 3
	} else {
		return 2*f(n-1) + 1
	}
}

func main() {
	fmt.Println(f(5))
}
