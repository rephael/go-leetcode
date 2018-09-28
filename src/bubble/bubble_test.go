package bubble

import (
	"fmt"
	"testing"
)

func Test_bubbleUp(t *testing.T) {
	array := []int{3, 2, 5, 4, 7, 1, 9, 8}
	result := bubbleUp(array)
	fmt.Println(result)
}

func Test_swap(t *testing.T) {
	a := 1
	b := 2
	swap(&a, &b)
	fmt.Println(a, b)
}

func Test_bubbleSort(t *testing.T) {
	array := []int{3, 2, 5, 4, 7, 1, 9, 8}
	bubbleSort(array, len(array))
	fmt.Println(array)
}
