package main

import "fmt"

//二维数组的遍历

func main() {
	arr1 := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr1[i]); j++ {
			fmt.Printf("%v\t", arr1[i][j])
		}
		fmt.Println()
	}

	for i, v := range arr1 {
		for j, v1 := range v {
			fmt.Printf("arr1[%v][%v]:%v \t", i, j, v1)
		}
		fmt.Println()
	}
}
