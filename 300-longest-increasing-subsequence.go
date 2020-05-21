package leetcode

import (
	"fmt"
)

/**
给定一个无序的整数数组，找到其中最长上升子序列的长度。

示例:

输入: [10,9,2,5,3,7,101,18]
输出: 4
解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。
说明:

可能会有多种最长上升子序列的组合，你只需要输出对应的长度即可。
你算法的时间复杂度应该为 O(n2) 。
进阶: 你能将算法的时间复杂度降低到 O(n log n) 吗?

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-increasing-subsequence
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
最值问题: 动态规划;
	状态转移方程 第n位最大上升长度为dp(n)
	dp[i] 为考虑前 i 个元素，以第 i 个数字结尾的最长上升子序列的长度
	dp[i]=max(dp[j])+1,其中0≤j<i且num[j]<num[i]
*/

// dp[i]=max(dp[j])+1,其中0≤j<i且num[j]<num[i]
func lengthOfLIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	dp := []int{}
	maxLenth := 1
	for i := 0; i < len(nums); i++ {
		dp = append(dp, 1)
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
			fmt.Printf("i:%d j:%d, index:%d dp: %v \n", i, j, dp[i], dp)
		}
		if maxLenth < dp[i] {
			maxLenth = dp[i]
		}
		//fmt.Print(i, dp, "\n")
	}
	return maxLenth
}

/**
LIS 算法: 最大上升子序列
O(n * Log(n)) 二分查找
// https://leetcode-cn.com/problems/longest-increasing-subsequence/solution/dong-tai-gui-hua-she-ji-fang-fa-zhi-pai-you-xi-jia/
*/
// ^int(^uint(0) >> 1) int 最小值
// unit(0)       	   == 00 00 00 00
// ^uint(0)      	   == FF FF FF FF
// ^uint(0) >> 1 	   == 7F FF FF FF
// int(^uint(0) >> 1)  == 7F FF FF FF
// ^int(^uint(0) >> 1) == 10 00 00 00 -> 最小int
func lengthOfLIS2(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	q := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		if val := nums[i]; val > q[len(q)-1] {
			q = append(q, val)
		} else {
			q[lengthOfLISQuery(q, val)] = val
		}
	}
	return len(q)
}

func lengthOfLISQuery(arr []int, val int) int {
	start, end := 0, len(arr)
	for start < end {
		m := start + (end-start)/2
		if arr[m] < val {
			start = m + 1
		} else {
			end = m
		}
	}
	return start
}
