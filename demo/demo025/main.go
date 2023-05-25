package main

import "fmt"

func cal(n1 float64, n2 float64, op byte) float64 {
	var res float64
	switch op {
	case '+':
		res = n1 + n2
	case '-':
		res = n1 - n2
	case '*':
		res = n1 * n2
	case '/':
		res = n1 / n2
	}
	return res
}
func main() {
	resturl := cal(1.0, 2.0, '+')
	fmt.Println("resturl=", resturl)
}
