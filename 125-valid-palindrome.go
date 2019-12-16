package leetcode

import (
	"unicode"
)

/**
给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。

说明：本题中，我们将空字符串定义为有效的回文串。

示例 1:

输入: "A man, a plan, a canal: Panama"
输出: true
示例 2:

输入: "race a car"
输出: false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-palindrome
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
双指针处理
*/
func isPalindrome(s string) bool {
	if s == "" {
		return true
	}
	left := 0
	right := len(s) - 1
	for left < right {
		if !isValidLetterOrDigit(rune(s[left])) {
			left++
			continue
		}
		if !isValidLetterOrDigit(rune(s[right])) {
			right--
			continue
		}
		if unicode.ToLower(rune(s[left])) != unicode.ToLower(rune(s[right])) {
			return false
		}
		left++
		right--
	}
	return true
}

// isValidLetterOrDigit 验证是字母或者数字
func isValidLetterOrDigit(c rune) bool {
	if unicode.IsLetter(c) || unicode.IsDigit(c) {
		return true
	}
	return false
}
