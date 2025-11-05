package leetcode

import "sort"

// 一旦根据特征进行归类，使用hash 散列表
func groupAnagrams(strs []string) [][]string {

	groups := make(map[string][]string)	

	for _, str := range strs {
		bytes := []byte(str)
		sort.Slice(bytes, func(i, j int) bool {
			return bytes[i] < bytes[j]
		})

		key := string(bytes)
		groups[key] = append(groups[key], str)
	}

	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}

    return result
}