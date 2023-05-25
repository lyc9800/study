package main

import "fmt"

func main() {
	//统计3个班级的成绩情况，每个班级5个同学，求出各个班的平均分和所有班级的平均分
	//第一个版本固定为一个班级，5个学生
	//第二个版本，固定为3个班级，每个班级5个学生
	//定义班级个数，学生个数为变量
	//
	var classNum int = 3
	var stuNum int = 5
	var totalSum float32 = 0.0
	var passCount int = 0
	for j := 1; j <= classNum; j++ {
		var sum float32
		for i := 1; i <= stuNum; i++ {
			var score float32
			fmt.Printf("请输入%d班第%d个学生成绩\n", j, i)
			fmt.Scanln(&score)
			sum += score
			if score >= 60 {
				passCount++
			}
		}
		fmt.Printf("%v班总分为:%f 平均分为: %f\n", j, sum, sum/float32(stuNum))
		totalSum += sum
	}
	fmt.Printf("%d个班级总成绩是:%v %d个班级的平均成绩为: %v\n", classNum, totalSum, classNum, totalSum/float32(classNum*stuNum))
	fmt.Println("及格人数为: ", passCount)
}
