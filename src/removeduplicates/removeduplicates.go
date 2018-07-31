package removeduplicates

// [1, 1, 2]
func removeDuplicates(nums []int) int {
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != nums[index] {
			index++
			nums[index] = nums[i]
		}
	}
	return index + 1
}
