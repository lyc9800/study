package main

import "fmt"

func main() {
	//找出0-100中间的奇数
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}
