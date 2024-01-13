package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 1, 2, 3, 4, 5, 1, 2, 2, 3, 3, 3}
	res := numIdenticalPairs(nums)
	fmt.Println(res)
}

func numIdenticalPairs(nums []int) int {
	ans := 0
	count := make(map[int]int)
	for _, value := range nums {
		ans += count[value]
		count[value]++

	}
	return ans
}
