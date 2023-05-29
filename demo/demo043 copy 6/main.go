//景区门票示例

package main

import "fmt"

type Visitor struct {
	Name string
	Age  int
}

func (visitor *Visitor) showPrice() {
	if visitor.Age >= 90 || visitor.Age < 18 {
		fmt.Println("考虑到安全，建议不要完了")
		return
	}
	if visitor.Age > 18 && visitor.Age < 90 {
		fmt.Printf("游客姓名: %v\n年龄：%v,请付费20元\n", visitor.Name, visitor.Age)
	} else {
		fmt.Printf("游客姓名: %v\n年龄：%v,您可以免费游玩\n", visitor.Name, visitor.Age)
	}
}


func main() {
	var v Visitor
	for {
		fmt.Println("请输入你的名字:")
		fmt.Scanln(&v.Name)
		if v.Name == "n" {
			fmt.Println("程序退出……")
			break
		}
		fmt.Println("请输入你的年龄：")
		fmt.Scanln(&v.Age)
		v.showPrice()
	}
}
