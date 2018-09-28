package bubble

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func bubbleUp(array []int) []int {
	length := len(array)
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-1-i; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}

func bubbleSort(array []int, length int) {
	if length < 2 {
		return
	}
	for index := range array {
		if index > len(array)-2 {
			break
		}
		if array[index] > array[index+1] {
			array[index], array[index+1] = array[index+1], array[index]
		}
	}
	bubbleSort(array, length-1)
}
