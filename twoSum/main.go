package main

import "fmt"

func main() {
	nums := []int{3, 2, 5, 5, 5, 5, 3, 0}
	target := 6
	result := twoSum(nums, target)
	fmt.Println(result)
}

// this function checks if there are solutions that include duplicate numbers
// like [3,3] with target=6
// this is needed because hashmap doesn't memorize duplicate values of input array
func checkDuplicates(nums []int, target int) []int {
	if target%2 == 0 {
		result := []int{}
		tmp := target / 2
		for i := range nums {
			if nums[i] == tmp {
				result = append(result, i)
			}
		}
		if len(result) == 2 {
			return result
		}
	}
	return []int{}
}

func twoSum(nums []int, target int) []int {
	//check if duplicate solution exists
	result := checkDuplicates(nums, target)
	if len(result) != 0 {
		return result
	}
	//if no duplicate solution, go with hmap
	hmap := make(map[int]int)

	for index, num := range nums {
		hmap[num] = index
	}

	quicksort(nums, 0, len(nums)-1)

	//two-pointer approach
	leftPointer := 0
	rightPointer := len(nums) - 1

	for leftPointer < rightPointer {
		sum := nums[leftPointer] + nums[rightPointer]

		if sum > target {
			rightPointer--
		} else if sum < target {
			leftPointer++
		} else {
			return []int{hmap[nums[leftPointer]], hmap[nums[rightPointer]]}
		}
	}

	return []int{}
}

func quicksort(arr []int, leftIndex, rightIndex int) {
	if leftIndex < rightIndex {
		pivotIndex := partition(arr, leftIndex, rightIndex)

		quicksort(arr, leftIndex, pivotIndex-1)
		quicksort(arr, pivotIndex+1, rightIndex)
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
