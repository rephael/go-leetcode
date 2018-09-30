package selectionsort

import (
	"fmt"
)

// 选择排序普通实现
func selectSort(array []int) {
	for i := 0; i < len(array)-1; i++ {
		fmt.Println(array)
		minIndex := i
		for j := i + 1; j < len(array); j++ {
			if array[minIndex] > array[j] {
				minIndex = j
			}
		}
		if minIndex != i {
			array[minIndex], array[i] = array[i], array[minIndex]
		}
	}
}

// 找到数组元素中的最小值下标
func findMinArray(array []int) int {
	var minIndex int
	for i := 0; i < len(array) - 1; i++ {
		if array[minIndex] > array[i] {
			minIndex = i
		}
	}
	return minIndex;
}