package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

type question128 struct {
	para128
	ans128
}

// para 是参数
// one 代表第一个参数
type para128 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans128 struct {
	one int
}

func Test_Problem128(t *testing.T) {

	qs := []question128{

		{
			para128{[]int{}},
			ans128{0},
		},

		{
			para128{[]int{0}},
			ans128{1},
		},

		{
			para128{[]int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6}},
			ans128{7},
		},

		{
			para128{[]int{2147483646, -2147483647, 0, 2, 2147483644, -2147483645, 2147483645}},
			ans128{3},
		},

		{
			para128{[]int{100, 4, 200, 1, 3, 2}},
			ans128{4},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 128------------------------\n")

	for i, q := range qs {
		expected := q.ans128.one
		actual := longestConsecutive(q.para128.one)

		// 使用 reflect.DeepEqual 进行深度比较（处理切片、数组、map等复杂类型）
		// 对于空切片，normalize 空切片和 nil 的比较
		aNorm, eNorm := actual, expected

		pass := reflect.DeepEqual(aNorm, eNorm)
		if pass {
			fmt.Printf("【测试用例 %d】✅ PASS\n", i+1)
		} else {
			fmt.Printf("【测试用例 %d】❌ FAIL\n", i+1)
			t.Errorf("测试用例 %d 失败: 期望 %v, 实际得到 %v", i+1, expected, actual)
		}
		fmt.Printf("  【input】:%v\n", q.para128)
		fmt.Printf("  【expected】:%v\n", expected)
		fmt.Printf("  【actual】:%v\n", actual)
	}
	fmt.Printf("\n\n\n")
}