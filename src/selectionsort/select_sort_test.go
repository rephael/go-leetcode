package selectionsort

import (
	"fmt"
	"testing"
)

func Test_select_sort(t *testing.T) {
	array := []int{3, 2, 5, 4, 7, 1, 9, 8, 6,10,14,12,13}
	selectSort(array)
	fmt.Println(array)
}

func Test_findMinArray(t *testing.T) {
	array := []int{3, 2, 5, 4, 7, 1, 9, 8, 6,10,14,12,13}
	fmt.Println(findMinArray(array))
}