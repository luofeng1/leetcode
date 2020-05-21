package leetcode

import (
	"sort"
)

/**
给定一些标记了宽度和高度的信封，宽度和高度以整数对形式 (w, h) 出现。当另一个信封的宽度和高度都比这个信封大的时候，这个信封就可以放进另一个信封里，如同俄罗斯套娃一样。

请计算最多能有多少个信封能组成一组“俄罗斯套娃”信封（即可以把一个信封放到另一个信封里面）。

说明:
不允许旋转信封。

示例:

输入: envelopes = [[5,4],[6,4],[6,7],[2,3]]
输出: 3
解释: 最多信封的个数为 3, 组合为: [2,3] => [5,4] => [6,7]。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/russian-doll-envelopes
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func maxEnvelopes(envelopes [][]int) int {
	if len(envelopes) <= 1 {
		return len(envelopes)
	}
	sort.Sort(maxEnvelopesSort(envelopes))

	maxEnvelopes := 1
	dp := make([]int, len(envelopes), len(envelopes))
	dp[0] = 1
	for i := 1; i < len(envelopes); i++ {
		width := envelopes[i][1]
		length := envelopes[i][0]
		currentMaxEnvelopes := 1
		for j := 0; j < i; j++ {
			width2 := envelopes[j][1]
			length2 := envelopes[j][0]
			if width > width2 && length > length2 && dp[j]+1 > currentMaxEnvelopes {
				//fmt.Printf("item:%v j:item:%v current:%d <> %d %d %d %d\n", envelopes[i], envelopes[j], currentMaxEnvelopes, dp[j], currentMaxEnvelopes, width, width2)
				currentMaxEnvelopes = dp[j] + 1
			}
		}
		dp[i] += currentMaxEnvelopes
		//fmt.Printf("--item:%v max:%d \n", envelopes[i], dp[i])
		if dp[i] > maxEnvelopes {
			maxEnvelopes = dp[i]
		}
	}
	return maxEnvelopes
}

type maxEnvelopesSort [][]int

//Len()
func (s maxEnvelopesSort) Len() int {
	return len(s)
}

//Less(): 成绩将有低到高排序
func (s maxEnvelopesSort) Less(i, j int) bool {
	return s[i][0] < s[j][0]
}

// Swap()
func (s maxEnvelopesSort) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

//func maxEnvelopes(envelopes [][]int) int {
//	sort.Slice(envelopes, func(i, j int) bool {
//		if envelopes[i][0] == envelopes[j][0] {
//			return envelopes[i][1] < envelopes[j][1]
//		}
//		return envelopes[i][0] > envelopes[j][0]
//	})
//	k := 0
//	for _, v := range envelopes {
//		j := sort.Search(k, func(i int) bool {
//			c := envelopes[i]
//			return c[0] <= v[0] || c[1] <= v[1]
//		})
//		envelopes[j] = v
//		if j == k {
//			k++
//		}
//	}
//	return k
//}
