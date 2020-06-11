package leetcode

import (
	"bytes"
	"strconv"
)

/**
给定一个经过编码的字符串，返回它解码后的字符串。

编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。

你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。

此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。

示例 1：

输入：s = "3[a]2[bc]"
输出："aaabcbc"
示例 2：

输入：s = "3[a2[c]]"
输出："accaccacc"
示例 3：

输入：s = "2[abc]3[cd]ef"
输出："abcabccdcdcdef"
示例 4：

输入：s = "abc3[cd]xyz"
输出："abccdcdcdxyz"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/decode-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func decodeString(s string) string {
	position := 0
	return string(decodeStringDFS([]byte(s), &position))
}

func decodeStringDFS(s []byte, position *int) []byte {
	b := bytes.Buffer{}
	num := bytes.Buffer{}
	for *position < len(s) {
		// ] 结束则返回 || 最后一位则返回
		if *position >= len(s) || s[*position] == byte(']') {
			*position++
			return b.Bytes()
		}

		// 数字则合成数字
		if s[*position] >= byte('0') && s[*position] <= byte('9') {
			num.WriteByte(s[*position])
			*position++
			continue
		}

		// [ 开始递归
		if s[*position] == byte('[') {
			*position++
			repeatBytes := decodeStringDFS(s, position)
			for i, _ := strconv.Atoi(string(num.Bytes())); i > 0; i-- {
				b.WriteString(string(repeatBytes))
			}
			num = bytes.Buffer{}
			continue
		}

		// 字母则拼接
		b.WriteByte(s[*position])
		*position ++
	}

	return b.Bytes()
}
