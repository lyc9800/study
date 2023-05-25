package main

import "fmt"

func test() bool {
	fmt.Println("test……")
	return true
}
func main() {
	var i int = 10
	if i < 9 && test() {
		fmt.Println("这个调用了test的函数")
	} else {
		fmt.Println("这个没有调用test的函数")
	}

}
