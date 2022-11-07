package main

import "fmt"

func main() {
	//转义字符演示
	fmt.Println("Hello\tWorld")    //  \t表示制表符
	fmt.Println("Hello\nWorld")    //  \n表示换行
	fmt.Println("Hello\\World")    //  \\表示\
	fmt.Println("Hello \"World\"") // \"表示引号
	fmt.Print("你好世界\r拥抱\n")          //

	fmt.Println("姓名\t年龄\t籍贯\t住址\n张三\t12\t河北\t北京")
}
