package leetcode


func twoSum(nums []int, target int) []int {

	m := make(map[int]int)
	for k, v := range nums {
		if idx, ok := m[target-1]; ok {
			return []int{idx, k}
		}
		m[v] = k
	}
	return nil	
}