package main

import "fmt"

var name = "Tom"

func test01() {
	fmt.Println("test01  name=", name)
}

func test02() {
	name := "Jack"
	fmt.Println("test02  name=", name)
}

func main() {
	fmt.Println("name=", name)
	test01()
	test02()
	test01()
}
