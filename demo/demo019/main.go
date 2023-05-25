package main

import "fmt"

func main() {
	//打印金字塔
	//1，先写死打印出来一个矩形
	//第二次修改让j小于i
	//第三次修改，增加空格的输入
	//第四次修改，增加前面的空格
	//第五次修改，让用户输入层数
	var layers int
	fmt.Println("请输入金字塔的层数，最小为3")
	fmt.Scanln(&layers)

	for a := 1; a <= layers; a++ {
		for k := 1; k <= layers-a; k++ {
			fmt.Print(" ")
		}

		for b := 1; b <= 2*a-1; b++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}

}
