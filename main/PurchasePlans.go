package main

import (
	"fmt"
	"sort"
)

func PurchasePlans(nums []int, target int) int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	for _, num := range nums {
		fmt.Println(num)
	}
	fmt.Println(target)
	return 0
}
