package main

import "fmt"

func main() {
	var days int = 97
	//var weeks int = days / 7
	fmt.Printf("总共有 %v 周\n", days/7)
	fmt.Printf("还剩下 %v 天\n", days%7)
}
