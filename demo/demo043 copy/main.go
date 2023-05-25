package main

import (
    "fmt"
    "math/rand"
    "sort"
)

func main() {
    // 生成10个随机整数
    nums := make([]int, 10)
    for i := 0; i < 10; i++ {
        nums[i] = rand.Intn(100) + 1
    }
    fmt.Println("随机生成的10个整数：", nums)

    // 冒泡法排序
    for i := 0; i < len(nums)-1; i++ {
        for j := 0; j < len(nums)-1-i; j++ {
            if nums[j] > nums[j+1] {
                nums[j], nums[j+1] = nums[j+1], nums[j]
            }
        }
    }
    fmt.Println("排序后的整数：", nums)

    // 二分法查找90
    index := sort.SearchInts(nums, 90)
    if index < len(nums) && nums[index] == 90 {
        fmt.Println("数字90的下标为：", index)
    } else {
        fmt.Println("没有查找到数字90")
    }
}
