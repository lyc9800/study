package main

import (
	"fmt"
)

func main() {
	var intArr [5]int = [...]int{11, 22, 33, 44, 55}
	fmt.Println(intArr)

	slice := intArr[1:3]

	fmt.Println("slice的内容是:", slice)
	fmt.Println("slice的长度是:", len(slice))
	fmt.Println("slice的容量是:", cap(slice))

	/* for i:=0;i<len(slice);i++{
		fmt.Printf("slice的地址是:%p\n", &slice[i])
	} */

	fmt.Println("intArr第一个元素的地址是:",&intArr[1])
	fmt.Println("slice的第一个元素地址是:",&slice[0])

}
