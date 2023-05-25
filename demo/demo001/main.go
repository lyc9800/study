package main

import (
	"fmt"
)

func main() {
	var i int = 10
	fmt.Println("i的值=", i, "i的内存地址是:", &i)

	var ptr *int
	ptr = &i
	*ptr = 100
	fmt.Printf("i is=%v", i)
	fmt.Println("Ptr=", *ptr)
}
