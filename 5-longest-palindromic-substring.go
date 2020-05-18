package leetcode

import (
	"strings"
)

/**
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
示例 2：

输入: "cbbd"
输出: "bb"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-palindromic-substring
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
动态规划: 多阶段决策最优解模型
状态转移方程: dp[i][j]

最优子结构: 后面阶段的状态可以通过前面阶段的状态推导出来
无后效性:   前面阶段的状态确定之后，不会被后面阶段的决策所改变
重复子问题: 不同的决策序列，到达某个相同的阶段时，可能会产生重复的状态

*/
/**
解决这类问题的核心思想就是两个字“延伸”，具体来说

如果一个字符串是回文串，那么在它左右分别加上一个相同的字符，那么它一定还是一个回文串
如果一个字符串不是回文串，或者在回文串左右分别加不同的字符，得到的一定不是回文串
事实上，上面的分析已经建立了大问题和小问题之间的关联， 基于此，我们可以建立动态规划模型。

我们可以用 dp[i][j] 表示 s 中从 i 到 j（包括 i 和 j）是否可以形成回文， 状态转移方程只是将上面的描述转化为代码即可：
if (s[i] === s[j] && dp[i + 1][j - 1]) {
  dp[i][j] = true;
}
P(i,j)=(P(i+1,j−1)&&S[i]==S[j])
*/
func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	sArr := strings.Split(s, "")
	maxStr := ""
	maxAmount := 0
	record := map[int]map[int]bool{} // 注意: map性能比数组性能差了很多;此处用map会导致超时
	for i := len(sArr) - 1; i >= 0; i-- {
		for j := i; j < len(sArr); j++ {
			isPalindromic := longestPalindromeIsPalindromic(&sArr, i, j, record)
			if _, exists := record[i]; !exists {
				record[i] = map[int]bool{}
			}
			record[i][j] = isPalindromic
			if maxAmount < j-i+1 && isPalindromic {
				maxStr = strings.Join(sArr[i:j+1], "")
				maxAmount = j - i + 1
			}
		}
	}
	return maxStr
}

func longestPalindromeIsPalindromic(arr *[]string, i, j int, record map[int]map[int]bool) bool {
	if i == j || (i == j-1 && (*arr)[i] == (*arr)[j]) {
		return true
	}
	if _, exists := record[i+1]; exists {
		if v, exists := record[i+1][j-1]; exists && v && (*arr)[i] == (*arr)[j] {
			return true
		}
	}
	return false
}

/**
所以，当我们拿到问题的时候，我们可以先用简单的回溯算法解决，然后定义状态，每个状态表示一个节点，然后对应画出递归树。从递归树中，我们很容易可以看出来，是否存在重复子问题，以及重复子问题是如何产生的。以此来寻找规律，看是否能用动态规划解决。

babad
1. 穷举:
f(0,0) f(0,1) f(0,2) f(0, 3)
f(1,1) f(1,2) f(1,3)
f(2,2) f(2,3) f(2,4)
f(3,3) f(3,4)
f(4,4)
定义状态: 是不是回文字符串; dp(i,j) 代表是不是回文字符串
dp(i,j) Y => arr[i-1]==arr[j+1] && dp(i-1,j+1)
i == j => dp(i,j) => Y
i == j-1 && arr[i-1]==arr[j+1] => Y
数组:
*/

/**
内存占用极高
*/
func longestPalindrome2(s string) string {
	if s == "" {
		return ""
	}
	if len(s) == 1 {
		return s
	}
	if len(s) == 2 {
		if s[0] == s[1] {
			return s
		}
		return s[0:1]
	}
	const (
		unStar = iota
		true
		false
	)
	dp := make([][]int, len(s), len(s))
	var maxIndex int = 1
	var maxArr = [2]int{0, 1}
	for i := 0; i < len(s); i++ {
		dp[i] = make([]int, len(s), len(s))
		dp[i][i] = true
	}
	for i := 1; i < len(s); i++ {
		for j := 0; j < len(s)-i; j++ {
			dp[j][j+i] = false
			if (j == j+i-1 || dp[j+1][j+i-1] == true) && s[j] == s[j+i] {
				dp[j][j+i] = true
				if i+1 > maxIndex {
					maxIndex = i
					maxArr = [2]int{j, j + i + 1}
				}
			}
		}
	}
	return s[maxArr[0]:maxArr[1]]
}

/**
动态规划: 求最值 「自底向上」
	[穷举] 中 包含 [重叠子问题]
	一般包含 [最优子结构] 正确的 [状态转移方程]
	明确「状态」 -> 定义 dp 数组/函数的含义 -> 明确「选择」-> 明确 base case。
	先暴力 -> 确定其含有重叠子问题
*/

/**
动态规划: 凑零钱问题;
先看下题目：给你k种面值的硬币，面值分别为c1, c2 ... ck，
每种硬币的数量无限，再给一个总金额amount，问你最少需要几枚硬币凑出这个金额，
如果不可能凑出，算法返回 -1 。
// coins 中是可选硬币面值，amount 是目标金额
int coinChange(int[] coins, int amount);
*/

/**
暴力方法:
1元 2元 5元
凑成 11元
1
2
5
1 2
1 5
2 5
1 2 5

确定状态: 也就是原问题和子问题中变化的变量
由于硬币数量无限，所以唯一的状态就是目标金额amount。
然后确定dp函数的定义：函数 dp(n)表示，当前的目标金额是n，至少需要dp(n)个硬币凑出该金额。

*/

/**
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
示例 1：
输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
示例 2：
输入: "cbbd"
输出: "bb"

确定状态: 变化: 谁开始,谁结束, 回文长度 dp(start, end) = Y/N
bab
1-1 shi
1-2
1-3
2-2 shi
2-3
3-3 shi
[
start = end len = 1
end > start:
{
	dp(start, end) = dp(start+1, end-1) != 0 ?
		dp(start+1, end-1) + (str[start] === str[end] ? 2 : 0):
		0;
}
]
*/

/**
间隔, 规律
*/
func longestPalindrome3(s string) string {
	if s == "" {
		return ""
	}
	if len(s) == 1 {
		return s
	}
	if len(s) == 2 {
		if s[0] == s[1] {
			return s
		}
		return s[0:1]
	}
	lenMax := 1
	start := 0
	end := 0
	dp := make([][]bool, len(s), len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s), len(s))
		dp[i][i] = true
		if i+1 < len(s) {
			if s[i] == s[i+1] {
				dp[i][i+1] = true
				lenMax = 2
				start = i
				end = i + 1
			} else {
				dp[i][i+1] = false
			}
		}
		//fmt.Printf("i:%d max:%v \n", i, dp[i])
	}
	for j := 2; j < len(s); j++ {
		for i := 0; i <= j-2; i++ {
			if dp[i+1][j-1] && s[i] == s[j] {
				dp[i][j] = true
				if newMax := j - i + 1; newMax > lenMax {
					lenMax = newMax
					start = i
					end = j
				}
			} else {
				dp[i][j] = false
			}
			//fmt.Printf("i:%d j:%d max:%t pre:%t i:%s j:%s lenMax:%d \n", i, j, dp[i][j], dp[i+1][j-1], string(s[i]), string(s[j]), lenMax)
		}
	}
	return s[start : end+1]
}
