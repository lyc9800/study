// 流程控制
package main

import "fmt"

func main() {
	//条件控制
	var age int
	fmt.Print("please input your age:")
	fmt.Scan(&age)
	if age >18 {
		fmt.Printf("your age is %d,adult\n",age)
	}else if age <18{
		fmt.Printf("your age is %d,teenager\n",age)
	}

}
