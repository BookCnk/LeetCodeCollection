package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums) // Step 1: Sort the array

	for i := 0; i < len(nums)-2; i++ {
		// Step 4: Avoid duplicates for the fixed element
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// Two pointers
		left, right := i+1, len(nums)-1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				// Step 4: Avoid duplicates for the left and right pointers
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return result
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums)) // Expected output: [[-1, -1, 2], [-1, 0, 1]]
}
