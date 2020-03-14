package leetcode

import "fmt"

/**
给定一个非负整数数组，a1, a2, ..., an, 和一个目标数，S。现在你有两个符号 + 和 -。对于数组中的任意一个整数，你都可以从 + 或 -中选择一个符号添加在前面。

返回可以使最终数组和为目标数 S 的所有添加符号的方法数。

示例 1:

输入: nums: [1, 1, 1, 1, 1], S: 3
输出: 5
解释:

-1+1+1+1+1 = 3
+1-1+1+1+1 = 3
+1+1-1+1+1 = 3
+1+1+1-1+1 = 3
+1+1+1+1-1 = 3

一共有5种方法让最终目标和为3。
注意:

数组非空，且长度不会超过20。
初始的数组的和不会超过1000。
保证返回的最终结果能被32位整数存下。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/target-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 暴力解法
func findTargetSumWays(nums []int, S int) int {
	if len(nums) == 0 {
		return 0
	}
	var result int

	dp := []int{nums[0], nums[0] * -1}
	if (nums[0] == S || -1*nums[0] == S) && len(nums) == 1 {
		result = 1
	}
	for i := 1; i < len(nums); i++ {
		sum := + nums[i]
		dep := - nums[i]
		newDP := []int{}
		for _, v := range dp {
			if i == len(nums)-1 {
				if v+dep == S {
					result += 1
				}
				if v+sum == S {
					result += 1
				}
			}
			newDP = append(newDP, v+sum, v+dep)
		}
		dp = newDP
	}
	fmt.Println(dp)
	return result
}

/**
我们用 dp[i][j] 表示用数组中的前 i 个元素，组成和为 j 的方案数。考虑第 i 个数 nums[i]，它可以被添加 + 或 -，因此状态转移方程如下：

dp[i][j] = dp[i - 1][j - nums[i]] + dp[i - 1][j + nums[i]]

作者：LeetCode
链接：https://leetcode-cn.com/problems/target-sum/solution/mu-biao-he-by-leetcode/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

/**
该题难点在于，如何将其转换为背包问题。
2*sum(N) - sum(P) + sum(P) = S + sum(N) + sum(P)
=>
2*sum(N) = S + sum(nums)
*/
