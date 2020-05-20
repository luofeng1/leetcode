package leetcode

/**
给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。

相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。

 

例如，给定三角形：

[
     [2],
    [3,4],
   [6,5,7],
  [4,1,8,3]
]
自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。
说明：

如果你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题，那么你的算法会很加分。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/triangle
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
最值问题: 动态规划
状态 dp 表示 当前(x,y)位置最小的路径;
dp(x,y) = triangle[x][y] + min(dp[x+1][y], dp[x+1][y+1])
*/
func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	if len(triangle) == 1 {
		return triangle[0][0]
	}
	dp := make([][]int, len(triangle), len(triangle))
	for i := 0; i < len(triangle); i++ {
		dp[i] = make([]int, len(triangle[i]), len(triangle[i]))
		for j := 0; j < len(triangle[i]); j++ {
			dp[i][j] = triangle[i][j]
		}
	}
	for i := len(triangle) - 2; i >= 0; i-- {
		//fmt.Printf("第%d层 result: %v \n", i, dp[i])
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] += minimumTotalMin(dp[i+1][j], dp[i+1][j+1])
		}
	}
	return dp[0][0]
}

func minimumTotalMin(x, y int) int {
	if x < y {
		return x
	}
	return y
}
