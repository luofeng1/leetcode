package leetcode

/**
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。

问总共有多少条不同的路径？



例如，上图是一个7 x 3 的网格。有多少可能的路径？

说明：m 和 n 的值均不超过 100。

示例 1:

输入: m = 3, n = 2
输出: 3
解释:
从左上角开始，总共有 3 条路径可以到达右下角。
1. 向右 -> 向右 -> 向下
2. 向右 -> 向下 -> 向右
3. 向下 -> 向右 -> 向右
示例 2:

输入: m = 7, n = 3
输出: 28

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/unique-paths
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
状态转移方程:
状态 到 m,n 的可能方案次数
则:
dp[m][n] = dp[m-1][n] + dp[m][n-1]
dp[m][1] = 1
dp[1][n] = 1
*/
func uniquePaths2(m int, n int) int {
	dp := make([][]int, m+1, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1, n+1)
		dp[i][0] = 0
	}
	for i := 0; i < n+1; i++ {
		dp[0][i] = 0
	}
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if j == 1 {
				dp[i][j] = 1
				continue
			}
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m][n]
}

// dp[m][n] = dp[m-1][n] + dp[m][n-1]
func uniquePaths(m int, n int) int {
	dp := make([]int, m+n, m+n)
	dp[0] = 1
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			dp[j] = dp[j] + dp[j-1]
		}
	}
	return dp[n-1]
}
