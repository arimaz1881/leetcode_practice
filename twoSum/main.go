package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	target := 6
	result := twoSum(nums, target)
	fmt.Println(result)
}

func twoSum(nums []int, target int) []int {
	hmap := make(map[int]int)
	for index, num := range nums {
		hmap[num] = index
	}
	sort(nums, 0, len(nums)-1)
	result := make([]int, 0)
	val1, val2 := 0, 0

	leftPointer := 0
	rightPointer := len(nums) - 1

	for leftPointer < rightPointer {
		sum := nums[leftPointer] + nums[rightPointer]
		if sum > target {
			rightPointer--
		} else if sum < target {
			leftPointer++
		} else {
			val1 = nums[leftPointer]
			val2 = nums[rightPointer]
		}
	}

	result = append(result, hmap[val1])
	result = append(result, hmap[val2])

	return result
}

func sort(arr []int, leftIndex, rightIndex int) {
	if leftIndex < rightIndex {
		pivotIndex := partition(arr, leftIndex, rightIndex)

		sort(arr, leftIndex, pivotIndex-1)
		sort(arr, pivotIndex+1, rightIndex)
	}
}

func partition(arr []int, leftIndex, rightIndex int) int {
	pivot := arr[rightIndex]

	pointer := leftIndex - 1

	for j := leftIndex; j < rightIndex; j++ {
		if arr[j] <= pivot {
			pointer++
			arr[pointer], arr[j] = arr[j], arr[pointer]
		}
	}

	arr[pointer+1], arr[rightIndex] = arr[rightIndex], arr[pointer+1]

	return pointer + 1
}
