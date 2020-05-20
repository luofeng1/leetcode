package leetcode

/**
在一个由 0 和 1 组成的二维矩阵内，找到只包含 1 的最大正方形，并返回其面积。

示例:

输入:

1 0 1 0 0
1 0 1 1 1
1 1 1 1 1
1 0 0 1 0

输出: 4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximal-square
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
最值问题: 动态规划;
确定状态: dp(x,y) 表示 只包含1的最大面积为 dp[x][y]
状态转移方程: dp(i,j)=min(dp(i−1,j),dp(i−1,j−1),dp(i,j−1))+1
*/
func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	maxSquare := 0
	dp := make([][]int, len(matrix), len(matrix))
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[0]), len(matrix[0]))
		for j := 0; j < len(matrix[0]); j++ {
			dp[i][j] = int(matrix[i][j] - '0')
			if dp[i][j] == 1 {
				maxSquare = 1
			}
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			val := dp[i][j]
			if val == 0 {
				dp[i][j] = 0
				continue
			}
			dp[i][j] = 1 + maximalSquareMin(maximalSquareMin(dp[i][j-1], dp[i-1][j]), dp[i-1][j-1])
			//fmt.Printf("x:%d y:%d val:%d dp:%d \n", i, j, val, dp[i][j])
			if dp[i][j] > maxSquare {
				maxSquare = dp[i][j]
			}
		}
	}
	return maxSquare * maxSquare
}

func maximalSquareMin(x, y int) int {
	if x < y {
		return x
	}
	return y
}
