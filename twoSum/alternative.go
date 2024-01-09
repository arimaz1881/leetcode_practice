package main

import "fmt"

func main() {
	nums := []int{5, 2, 3, 5, 5, 3, 0}
	target := 6
	result := twoSum(nums, target)
	fmt.Println(result)
}


func twoSum(nums []int, target int) []int {
	hmap := make(map[int]int)
	tmp := 0
	exists := false
	for i := range nums {
		tmp = target - nums[i]
		_, exists = hmap[tmp]
		if exists == true {
			return []int{i, hmap[tmp]}
		}
		hmap[nums[i]] = i
	}
	return []int{}
}
