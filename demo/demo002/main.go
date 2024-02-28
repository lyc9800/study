// 基本数据类型转换string
// 使用strconv转换string
package main

import (
	"fmt"
	"strconv"
)

func main() {
	//基本数据类型int转换string
	var i int32 = 99
	var str string
	str = strconv.FormatInt(int64(i), 10)
	fmt.Printf("%q, %T\n", str, str)

	//bool转换string

	var str1 string = "true"
	var b bool
	b, _ = strconv.ParseBool(str1)
	fmt.Printf("%t, %T\n", b, b)

	var str2 string = "123456"
	var n1 int64
	n1, _ = strconv.ParseInt(str2, 10, 64)
	fmt.Printf("%v, %T\n", n1, n1)
}
