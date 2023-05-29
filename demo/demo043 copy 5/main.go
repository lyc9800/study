package main

import "fmt"

type Student struct {
	name   string
	gender string
	age    int
	id     int
	score  float64
}

func (Student *Student) say() string {
	infoStr := fmt.Sprintf("student的信息: name=[%v] gender=[%v],age=[%v] id=[%v] score=[%v]",
	Student.name,Student.gender,Student.age,Student.id,Student.score)
	 return infoStr
}

func main() {
	var stu = Student{
		name : "Luoyichao",
		gender : "男",
		age : 33,
		id : 100,
		score : 999,
	}
	fmt.Println(stu.say())
}