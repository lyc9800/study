package main

import "fmt"

func main() {
	//字母转化程序，将输入的小写字母换成大写字母，abcde  只能是这5个字母
	//其他字母输出other
	/* var key byte

	fmt.Println("请输入您要转化的字母(a,b,c,d,e)")
	fmt.Scanf("%c", &key)

	switch key {
	case 'a':
		fmt.Printf("%c转换后为:A", key)
	case 'b':
		fmt.Printf("%c转换后为:B", key)
	case 'c':
		fmt.Printf("%c转换后为:C", key)
	case 'd':
		fmt.Printf("%c转换后为:D", key)
	case 'e':
		fmt.Printf("%c转换后为:E", key)
	default:
		fmt.Printf("other")
	} */

	//根据输入的成绩判断是否合格
	/* var score float32
	fmt.Println("请输入您的成绩:")
	fmt.Scanln(&score)

	switch {
	case score >= 60 && score <= 100:
		fmt.Println("合格")
	default:
		fmt.Println("不合格")
	} */

	//根据输入的月份1-12判断是什么季节12-2：春季；3-5：夏季；6-87：秋季：9-11
	var month int8
	fmt.Println("请输入月份(1-12数字):")
	fmt.Scanln(&month)

	switch month {
	case 3, 4, 5:
		fmt.Println("该月份季节是‘春季’")
	case 6, 7, 8:
		fmt.Println("该月份季节是‘夏季’")
	case 9, 10, 11:
		fmt.Println("该月份季节是‘秋季’")
	case 12, 1, 2:
		fmt.Println("该月份季节是‘冬季’")
	default:
		fmt.Println("输入的月份无效")
	}
}
