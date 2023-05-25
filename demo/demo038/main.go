package main

import (
	"fmt"
	"time"
)

func main() {

	n1 := time.Date(2023, time.April, 1, 15, 52, 0, 0, time.Local)
	fmt.Println(n1)
	fmt.Printf("当前系统时间为: %d-%d-%d %d:%d:%d\n", time.Now().Year(),
		time.Now().Month(), time.Now().Day(), time.Now().Hour(),
		time.Now().Minute(), time.Now().Second())

	fmt.Println("当前的月份是", time.Now().Format("01"))
	fmt.Println("格式化输出当前的时间是:", time.Now().Format("2006-01-02 15:04:05"))

	n2 := time.Now().UnixNano()
	fmt.Println(n2)
}
