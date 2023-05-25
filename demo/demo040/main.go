package main

import "fmt"

func main() {
	var zimu [26]byte
	for i := 0; i < 26; i++ {
		zimu[i] = byte('A' + i)
		fmt.Printf("第%d个字母是:%c\n", i+1, zimu[i])
	}
}
