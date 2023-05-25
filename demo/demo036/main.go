package main

import "fmt"

func jinzita(layers int) {
	for a := 1; a <= layers; a++ {
		for k := 1; k <= layers-a; k++ {
			fmt.Print(" ")
		}

		for b := 1; b <= 2*a-1; b++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}
func main() {
	var layers int
	fmt.Println("请输入金字塔的层数:")
	fmt.Scanln(&layers)
	jinzita(layers)
}
