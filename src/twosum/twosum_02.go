package twosum

func twoSum2(nums []int, target int) []int {
	data := make(map[int]int)
	for index, num := range nums {
		data[num] = index
	}
	for index, num := range nums {
		toFind := target - num
		i, exsits := data[toFind]
		if exsits && i != index {
			return []int{index, i}
		}
	}
	return nil
}
