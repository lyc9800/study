package main

import "fmt"

func main() {
	//判断一个数字是不是大于0，小于0，等于0
	var num int
	fmt.Println("请输入一个数字:")
	fmt.Scanln(&num)

	if num > 0 {
		fmt.Println("这个数字大于0")
	} else if num < 0 {
		fmt.Println("这个数字小于0")
	} else {
		fmt.Println("这个数字等于0")
	}

	//判断输入的年份是不是闰年
	//1）四年一闰百年不闰：即如果year能够被4整除，但是不能被100整除，则year是闰年。
	//2）每四百年再一闰：如果year能够被400整除，则year是闰年。

	var years int
	fmt.Println("请输入一个年份:")
	fmt.Scanln(&years)

	if (years%4 == 0 && years%100 != 0) || years%400 == 0 {
		fmt.Println("这是一个闰年")
	} else {
		fmt.Println("这不是一个闰年")
	}

	//判断一个数字是不是水仙花数，一个三位数，各个位的立方之和等于三位数本身 135=1x1x1=3x3x3=5x5x5

	var (
		number int
		var1   int
		var2   int
		var3   int
		var4   int
	)

	fmt.Println("请输入一个三位数字:")
	fmt.Scanln(&number)

	var1 = number / 100
	var2 = (number - var1*100) / 10
	var3 = number - var1*100 - var2*10
	var4 = var1*var1*var1 + var2*var2*var2 + var3*var3*var3

	if var4 == number {
		fmt.Printf("%d 是一个水仙花数字", number)
	} else {
		fmt.Printf("%d 不是一个水仙花数字", number)
	}

}
