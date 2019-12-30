package leetcode

/**
你将获得 K 个鸡蛋，并可以使用一栋从 1 到 N  共有 N 层楼的建筑。

每个蛋的功能都是一样的，如果一个蛋碎了，你就不能再把它掉下去。

你知道存在楼层 F ，满足 0 <= F <= N 任何从高于 F 的楼层落下的鸡蛋都会碎，从 F 楼层或比它低的楼层落下的鸡蛋都不会破。

每次移动，你可以取一个鸡蛋（如果你有完整的鸡蛋）并把它从任一楼层 X 扔下（满足 1 <= X <= N）。

你的目标是确切地知道 F 的值是多少。

无论 F 的初始值如何，你确定 F 的值的最小移动次数是多少？

 

示例 1：

输入：K = 1, N = 2
输出：2
解释：
鸡蛋从 1 楼掉落。如果它碎了，我们肯定知道 F = 0 。
否则，鸡蛋从 2 楼掉落。如果它碎了，我们肯定知道 F = 1 。
如果它没碎，那么我们肯定知道 F = 2 。
因此，在最坏的情况下我们需要移动 2 次以确定 F 是多少。
示例 2：

输入：K = 2, N = 6
输出：3
示例 3：

输入：K = 3, N = 14
输出：4
 

提示：

1 <= K <= 100
1 <= N <= 10000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/super-egg-drop
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
dp[m][k] 表示,k个鸡蛋,丢m次,能够达到最大的楼层数
*/
func superEggDrop(K int, N int) int {
	dp := make([][]int, N+1, N+1)
	for i := range dp {
		dp[i] = make([]int, K+1, K+1)
	}
	var m int
	for dp[m][K] < N {
		m += 1
		for i := 1; i <= K; i++ {
			dp[m][i] = dp[m-1][i-1] + 1 + dp[m-1][i]
		}
	}
	return m
}
// 以上,简化:
func superEggDrop2(K int, N int) int {
	dp := make([]int, K+1)
	m := 0
	for dp[K] < N {
		m++
		for i:=K; i>0; i-- {
			dp[i] = dp[i] + dp[i-1]+1
		}
	}
	return m
}