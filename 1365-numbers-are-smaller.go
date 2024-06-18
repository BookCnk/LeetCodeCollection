package main

func smallerNumbersThanCurrent(nums []int) []int {

	result := make([]int, len(nums))

	// [8, 1, 2, 2, 3]
	for i := 0; i < len(nums); i++ {
		count := 0
		for j := 0; j < len(nums); j++ {
			if nums[j] < nums[i] {
				count++
			}
		}
		result[i] = count
	}
	return result
}

// Inner Loop
//8 > 8 (false, count remains 0)
//8 > 1 (true, count becomes 1)
//8 > 2 (true, count becomes 2)
//8 > 2 (true, count becomes 3)
//8 > 3 (true, count becomes 4)
