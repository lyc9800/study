package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var count int = 0
	for {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(100) + 1
		count++

		if n == 99 {
			break
		}
	}
	fmt.Printf("生成99一共用了%d次", count)
}
