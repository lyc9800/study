//获取数字计算一个长方体的体积
//体积=长*宽*高

package main

import "fmt"

type Box struct{
	len float64
	width float64
	height float64
}

func (box *Box) getLong() float64 {
	return box.len * box.width * box.height
}

func main() {
	var box Box
	box.len = 10.0
	box.width = 10.0
	box.height =10.0
	long := box.getLong()
	fmt.Printf("体积是：%.2f",long)
}