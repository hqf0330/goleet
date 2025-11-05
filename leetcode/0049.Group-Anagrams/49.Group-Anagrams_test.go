package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

type question49 struct {
	para49
	ans49
}

// para 是参数
// one 代表第一个参数
type para49 struct {
	one []string
}

// ans 是答案
// one 代表第一个答案
type ans49 struct {
	one [][]string
}

func Test_Problem49(t *testing.T) {

	qs := []question49{

		{
			para49{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}},
			ans49{[][]string{{"ate", "eat", "tea"}, {"nat", "tan"}, {"bat"}}},
		},

		{
			para49{[]string{""}},
			ans49{[][]string{{""}}},
		},

		{
			para49{[]string{"a"}},
			ans49{[][]string{{"a"}}},
		},

	}

	fmt.Printf("------------------------Leetcode Problem 49------------------------\n")
	fmt.Printf("------------------------Leetcode Problem 1------------------------\n")

	for i, q := range qs {
		expected := q.ans49.one
		actual := groupAnagrams(q.para49.one)

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
		fmt.Printf("  【input】:%v\n", q.para49)
		fmt.Printf("  【expected】:%v\n", expected)
		fmt.Printf("  【actual】:%v\n", actual)
	}
	fmt.Printf("\n\n\n")
}