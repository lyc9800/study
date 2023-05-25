package main

import "fmt"

func main() {
	studenMap := make(map[string]map[string]string)
	studenMap["Num01"] = make(map[string]string, 3)
	studenMap["Num01"]["Name"] = "Tom"
	studenMap["Num01"]["Sex"] = "男"
	studenMap["Num01"]["Address"] = "北京长安街~"

	studenMap["Num02"] = make(map[string]string, 3)
	studenMap["Num02"]["Name"] = "Mary"
	studenMap["Num02"]["Sex"] = "女"
	studenMap["Num02"]["Address"] = "西安市长安区~"

	// fmt.Println(studenMap)
    // fmt.Println()
    // fmt.Println(studenMap["Num01"]["Address"])
    for k,v := range studenMap {
        fmt.Printf("Key:%v Value:%v",k,v)
        fmt.Printf("\t\n")
    }
}