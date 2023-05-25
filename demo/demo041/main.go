package main

import "fmt"

func main() {
	var intArray [5]int = [...]int{1, -1, 9, 99, 11}
	var sum float64 = 0.0

	for _, val := range intArray {
		sum += float64(val)
	}
	fmt.Printf("总大小为:%v 平均值为:%v", sum, sum/float64(len(intArray)))
}
