package leetcode

/**
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

示例:

输入: [-2,1,-3,4,-1,2,1,-5,4],
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
进阶:

如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-subarray
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
分治法: n = max(left, middle, right)
时间复杂度： O(nlogn) - n 是数组长度
空间复杂度： O(1)
*/
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return ^int(^uint(0) >> 1)
	}
	if len(nums) == 1 {
		return nums[0]
	}
	midIndex := len(nums) / 2
	middleLeft := 0
	middleMaxLeft := 0
	middleRight := 0
	middleMaxRight := 0
	if midIndex-1 >= 0 {
		for i := midIndex - 1; i >= 0; i-- {
			middleLeft += nums[i]
			if middleLeft > middleMaxLeft {
				middleMaxLeft = middleLeft
			}
		}
	}
	if midIndex+1 <= len(nums)-1 {
		for i := midIndex + 1; i < len(nums); i++ {
			middleRight += nums[i]
			if middleRight > middleMaxRight {
				middleMaxRight = middleRight
			}
		}
	}
	maxLeft := maxSubArray(nums[0:midIndex])
	maxMiddle := nums[midIndex] + middleMaxLeft + middleMaxRight
	maxRight := maxSubArray(nums[midIndex+1 : len(nums)])
	//fmt.Println(maxLeft, maxMiddle, maxRight)
	return maxSubArraymax3(maxLeft, maxMiddle, maxRight)
}

func maxSubArraymax3(left, middle, right int) int {
	if left >= middle && left >= right {
		return left
	}
	if middle >= left && middle >= right {
		return middle
	}
	return right
}

/**
动态规划1:
决策树: 最大: 开始or不开始
0: -2, 0
1: -1, 1, 0
2: -4, -2, -3, 0
3: 0, 2, 1, 4, 0
4: -1, 1, 0, 3, -1, 0
5: 1, 3, 2, 5, 1, 2, 0
6: 2, 4, 3, 6, 2, 3, 1, 0
7: -3, -1, -2, 1, -3, -2, -4, -5, 0
8: 1, 3, 2, 5, 1, 2, 0, -1, 4, 0
*/
func maxSubArray2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := []int{nums[0]}
	maxSum := nums[0]
	for i := 1; i < len(nums); i++ {
		newdp := []int{}
		maxSum = getMaxSum(0, nums, i, &newdp, maxSum)
		for _, v := range dp {
			maxSum = getMaxSum(v, nums, i, &newdp, maxSum)
		}
		//fmt.Println(newdp)
		dp = newdp
	}
	return maxSum
}

func getMaxSum(v int, nums []int, i int, newdp *[]int, maxSum int) int {
	n := v + nums[i]
	*newdp = append(*newdp, n)
	return maxSubArray2max2(maxSum, n)
}

func maxSubArray2max2(maxSum, num int) int {
	if maxSum > num {
		return maxSum
	}
	return num
}

/**
动态规划2:
动态规划的难点在于找到状态转移方程

dp[i] - 表示到当前位置 i 的最大子序列和
状态转移方程为： dp[i] = max(dp[i-1] + nums[i], nums[i])
初始化：dp[0] = nums[0]
 */
func maxSubArray3(nums []int) int {
	maxSum := nums[0]
	for i := 1; i < len(nums); i++ {
		nums[i] = maxSubArray2max2(0, nums[i-1]) + nums[i]
		if nums[i] > maxSum {
			maxSum = nums[i]
		}
	}
	return maxSum
}
