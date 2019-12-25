package leetcode

import (
	"strings"
)

/**
实现 strStr() 函数。

给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。

示例 1:

输入: haystack = "hello", needle = "ll"
输出: 2
示例 2:

输入: haystack = "aaaaa", needle = "bba"
输出: -1
说明:

当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。

对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与C语言的 strstr() 以及 Java的 indexOf() 定义相符。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/implement-strstr
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
使用KMP算法匹配
 */
func strStr(haystack, needle string) int {
	M := len(needle)
	if M == 0 {
		return 0
	}

	dp := NewKmp(needle)
	var j int
	for i := range haystack {
		j = dp[j][haystack[i]]
		// 到达终止态，返回结果
		if j == M {
			return i - M + 1
		}
	}
	return -1
}

// 简单粗暴有效....
func strStr2(haystack, needle string) int {
	return strings.Index(haystack, needle)
}

func NewKmp(pat string) [][]int {
	m := len(pat)
	dp := make([][]int, m, m)

	dp[0] = make([]int, 256, 256)
	dp[0][pat[0]] = 1

	var x int // 影子状态 X 初始为 0
	for j := 1; j < m; j++ {
		dp[j] = make([]int, 256, 256)
		for c := 0; c < 256; c++ {
			if int(pat[j]) == c {
				dp[j][c] = j + 1
			} else {
				dp[j][c] = dp[x][c]
			}
		}
		x = dp[x][pat[j]] // 更新影子状态
	}

	return dp
}

