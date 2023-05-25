package main

import "fmt"

func sum(n1, n2 float32) float32 {
	fmt.Printf("n1的类型是: %T\n", n1)
	return n1 + n2
}

func main() {
	fmt.Println("sum=", sum(1, 2))
}
