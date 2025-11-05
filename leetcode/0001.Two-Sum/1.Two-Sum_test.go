package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

type para struct {
	nums   []int
	target int
}

type ans struct {
	one []int
}

type question struct {
	para
	ans
}

func Test_Probelm(t *testing.T) {
	qs := []question{
		{
			para{[]int{3, 2, 4}, 6},
			ans{[]int{1, 2}},
		},

		{
			para{[]int{3, 2, 4}, 5},
			ans{[]int{0, 1}},
		},

		{
			para{[]int{0, 8, 7, 3, 3, 4, 2}, 11},
			ans{[]int{1, 3}},
		},

		{
			para{[]int{0, 1}, 1},
			ans{[]int{0, 1}},
		},

		{
			para{[]int{0, 3}, 5},
			ans{[]int{}},
		},
		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------Leetcode Problem 1------------------------\n")

	for i, q := range qs {
		expected := q.ans.one
		actual := twoSum(q.para.nums, q.para.target)

		// 使用 reflect.DeepEqual 进行深度比较（处理切片、数组、map等复杂类型）
		// 对于空切片，normalize 空切片和 nil 的比较
		aNorm, eNorm := actual, expected
		if len(actual) == 0 {
			aNorm = nil
		}
		if len(expected) == 0 {
			eNorm = nil
		}

		pass := reflect.DeepEqual(aNorm, eNorm)
		if pass {
			fmt.Printf("【测试用例 %d】✅ PASS\n", i+1)
		} else {
			fmt.Printf("【测试用例 %d】❌ FAIL\n", i+1)
			t.Errorf("测试用例 %d 失败: 期望 %v, 实际得到 %v", i+1, expected, actual)
		}
		fmt.Printf("  【input】:%v\n", q.para)
		fmt.Printf("  【expected】:%v\n", expected)
		fmt.Printf("  【actual】:%v\n", actual)
	}
	fmt.Printf("\n\n\n")
}
