package main

import "fmt"

func main() {
	//打印1~100之间所有是9的倍数的整数的个数及总和
	var conut int = 0
	var sum int = 0
	for i := 1; i <= 100; i++ {
		if i%9 == 0 {
			fmt.Printf("%v能被9整\n", i)
			sum += i
			conut++
		}
	}
	fmt.Printf("总共有%v个被9整除\n", conut)
	fmt.Println("能被9整除的总数之和=", sum)
}
