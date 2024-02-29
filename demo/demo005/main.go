// 练习
package main

import "fmt"

func main() {
	//假如还有97天放假，请问是几个星期几天
	days := 97
	weeks := days / 7
	remainder := days % 7
	fmt.Printf("还有%d天,也就是%d个星期%d天\n", days, weeks, remainder)

	//

}
