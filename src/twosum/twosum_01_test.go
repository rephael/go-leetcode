package twosum

import "testing"

func Test_twoSum1(t *testing.T) {
	nums := []int{2, 7, 9, 11}
	target := 9
	result := twoSum1(nums, target)
	if nums[result[0]]+nums[result[1]] == target {
		t.Log("OK")
	} else {
		t.Log("Failed")
	}
}
