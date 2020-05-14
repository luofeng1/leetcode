package leetcode

import (
	"fmt"
	"math"
)

/**
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:
abcadef
输入: "

"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
用一个hashmap来建立字符和其出现位置之间的映射。

维护一个滑动窗口，窗口内的都是没有重复的字符，去尽可能的扩大窗口的大小，窗口不停的向右滑动。

（1）如果当前遍历到的字符从未出现过，那么直接扩大右边界；

（2）如果当前遍历到的字符出现过，则缩小窗口（左边索引向右移动），然后继续观察当前遍历到的字符；

（3）重复（1）（2），直到左边索引无法再移动；

（4）维护一个结果res，每次用出现过的窗口大小来更新结果res，最后返回res获取结果。
*/

/**
也就是说，如果 s[j]s[j] 在 [i, j)[i,j) 范围内有与 j'j′重复的字符，我们不需要逐渐增加 ii 。
我们可以直接跳过 [i，j'][i，j′] 范围内的所有元素，并将 ii 变为 j' + 1j′+1。
*/
//func lengthOfLongestSubstring(s string) int {
//	if len(s) == 0 {
//		return 0
//	}
//	slidingWindow := map[uint8]int{s[0]: 0}
//	var start int
//	var end int
//	var max int = 1
//	for i := 1; i < len(s); i++ {
//		if v, exists := slidingWindow[s[i]]; exists && v >= start {
//			start = v + 1
//			slidingWindow[s[i]] = i
//			continue
//		}
//		end = i
//		max = int(math.Max(float64(max), float64(end-start+1)))
//		slidingWindow[s[i]] = i
//	}
//	return max
//}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	slideWindow := map[int32]int{}
	max := 1
	start := 0
	for k, v := range s {
		if val, exits := slideWindow[v]; !exits || val < start {
			slideWindow[v] = k
			max = int(math.Max(float64(max), float64(k-start+1)))
			fmt.Printf("max: %d start: %d k: %d v: %v slide: %v \n", max, start, k, string(rune(v)), slideWindow[v])
			continue
		}
		start = slideWindow[v] + 1
		slideWindow[v] = k
		fmt.Printf("unexits max: %d start: %d k: %d v: %v slide: %v \n", max, start, k, string(rune(v)), slideWindow[v])
	}
	fmt.Printf("finnal max: %d start: %d slide: %v \n", max, start, slideWindow)
	return max
}
