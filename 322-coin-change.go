package leetcode

/**
给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。

 

示例 1:

输入: coins = [1, 2, 5], amount = 11
输出: 3
解释: 11 = 5 + 5 + 1
示例 2:

输入: coins = [2], amount = 3
输出: -1
 

说明:
你可以认为每种硬币的数量是无限的。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/coin-change
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
变化状态: 金额n最少的数量是dp(n);
动态方程
dp(n) =
	0 n=0
	min(dp(n-[1,2,5])+1)
amount 11
*/
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	//dp := make(map[int]int, amount+1) // 慢..
	dp := make([]int, amount+1, amount+1) // 快 ✅
	for _, v := range coins {
		if v <= amount {
			dp[v] = 1
		}
	}
	//fmt.Println(0, dp)
	for i := 1; i <= amount; i++ {
		if dp[i] != 0 {
			for _, v := range coins {
				if i+v <= amount && (dp[i+v] == 0 || dp[i+v] > dp[i]+1) {
					dp[i+v] = dp[i] + 1
				}
			}
		}
	}
	//fmt.Println(dp[amount])
	if v := dp[amount]; v != 0 {
		return v
	}
	return -1
}
