package main

import "fmt"

func jiujiuchengfabiao() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %v\t", j, i, j*i)
		}
		fmt.Println()
	}

	for i := 9; i >= 1; i-- {
		for j := 9; j >= i; j-- {
			fmt.Printf("%v * %v = %v\t", j, i, j*i)
		}
		fmt.Println()
	}
}
func main() {
	jiujiuchengfabiao()
}
