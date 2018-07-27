package twosum

func twoSum3(nums []int, target int) []int {
	data := make(map[int]int)
	for index, num := range nums {
		toFind := target - num
		_, exsits := data[toFind]
		if exsits {
			return []int{data[toFind], index}
		}
		data[num] = index
	}
	return nil
}
