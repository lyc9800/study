package main

import (
	"fmt"
	"strconv"
)

func main() {
	//测试字符串的长度
	var str string = "hello 成都"
	fmt.Println("str的长度是:", len(str))

	//字符串转换
	str1 := []rune(str)
	for i := 0; i < len(str1); i++ {
		fmt.Printf("str1的字符是:%c\n", str1[i])
	}

	//字符串转整数，可以用来判断这个字符串是不是数字
	n, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("转换失败，失败原因是:", err)
	} else {
		fmt.Println("转换成功，转换后的值是:", n)
	}

	
}
