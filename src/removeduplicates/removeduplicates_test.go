package removeduplicates

import (
	"testing"
)

func Test_removeduplicates(t *testing.T) {
	nums := []int{1, 1, 2}
	length := removeDuplicates(nums)
	if length == 2 {
		t.Log("OK")
	} else {
		t.Log("Failed")
	}
}
