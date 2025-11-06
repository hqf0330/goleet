package leetcode

/*
一个无序数组，里面会有多种有序的片段，那么应该是说，一个片段一个kv放在散列表里
那么这个k应该怎么设计呢？value就是一个[]int就行了，最后把value的len输出好了



*/
func longestConsecutive(nums []int) int {


	numSet := make(map[int]bool)

	for _, num := range nums {
		numSet[num] = true	
	}

	longestStreak := 0
	for num := range numSet {
		if !numSet[num-1] {
			currentNum := num
			currentStreak := 1
			
			for numSet[currentNum+1] {
				currentNum++
				currentStreak++
			}

			if longestStreak < currentStreak {
				longestStreak = currentStreak
			}

		}
	}

    return longestStreak
}