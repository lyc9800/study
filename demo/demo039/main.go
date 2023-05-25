package main

import "fmt"

func main() {
	var intArray [3]int32
	for i := 0; i < len(intArray); i++ {
		fmt.Println("数组初始化的数值是:", intArray[i])
		fmt.Println("数组初始化的地址是:", &intArray[i])
	}
	fmt.Println()
	var scoreArray [5]float32
	var totilScorce float32

	for i := 0; i < len(scoreArray); i++ {
		fmt.Printf("请输入您的第%d个成绩:", i+1)
		fmt.Scanln(&scoreArray[i])
		totilScorce += scoreArray[i]
	}

	fmt.Println(totilScorce)
	for index, value := range scoreArray {
		fmt.Printf("第%d个成绩是:%.2f\n", index+1, value)
	}

}
