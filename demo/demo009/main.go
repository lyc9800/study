package main

import "fmt"

func main() {
	var i int = 123
	fmt.Printf("i的10进制是:%d\ni的8进制是:%0o\ni的16进制是:%0x", i, i, i)
}

// 011 100 101 345	
// 0 1110 0101 8+4+2 D5	
