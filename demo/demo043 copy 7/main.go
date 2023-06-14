package main

import "fmt"

type A struct {
	Name string
	Age  int
}

func (a *A) SayOk() {
	fmt.Println("A SayOk",a.Name)
}


func (a *A) hello(){
	fmt.Println("A hello",a.Name)
}

type B struct{
	A
}


func main() {
	var b B
	b.A.Name = "Tom"
	b.A.Age = 18
	b.A.SayOk()
	b.A.hello()

	b.Name = "Alex"
	b.Age = 20
	b.SayOk()
	b.hello()

}