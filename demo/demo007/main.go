package main

import "fmt"

func main() {
	/* var a int = 10
	var b int = 20
	t := a
	a = b
	b = t
	fmt.Printf("a的值是:%v\nb的值是:%v", a, b) */

	var n int = 100
	var m int = 200
	n = n + m
	m = n - m
	n = n - m
	fmt.Printf("n的数值是:%v\nm的数值是:%v", n, m)
}
