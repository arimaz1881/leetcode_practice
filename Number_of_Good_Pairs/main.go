package main

import "fmt"

func main() {
	nums := []int{5, 2, 3, 5, 5, 3, 0}
	res := numIdenticalPairs(nums)
	fmt.Println(res)
}

// прогнать фориком через массив и записывать в hmap, если элемент уже есть в мапе увеличивать каунт
func countUniq(nums []int) map[int]int {
	counts := make(map[int]int)
	for _, num := range nums {
		counts[num]++
	}

	return counts
}

// If a number appears n times, then n * (n – 1) // 2 good pairs can be made with this number.
func numIdenticalPairs(nums []int) int {
	counts := countUniq(nums)
	result := 0

	for _, count := range counts {
		result += count * (count - 1) / 2
	}

	return result
}
