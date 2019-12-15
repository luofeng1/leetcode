package leetcode

import "fmt"

/**
给定两个有序整数数组 nums1 和 nums2，将 nums2 合并到 nums1 中，使得 num1 成为一个有序数组。

说明:

初始化 nums1 和 nums2 的元素数量分别为 m 和 n。
你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
示例:

输入:
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3

输出: [1,2,2,3,5,6]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
我们需要三个指针：

current 用于记录当前填补到那个位置了

m 用于记录 nums1 数组处理到哪个元素了

n 用于记录 nums2 数组处理到哪个元素了
*/
func merge(nums1 []int, m int, nums2 []int, n int) []int {
	current := m + n - 1
	if n == 0 {
		return nums1
	}
	for i := current; i >= 0; i-- {
		fmt.Println(nums1, m, n, i)
		if n == 0 {
			return nums1
		}
		if m >= 1 && n >= 1 && nums1[m-1] > nums2[n-1] {
			nums1[i] = nums1[m-1]
			m -= 1
		} else {
			nums1[i] = nums2[n-1]
			n -= 1
		}
	}
	return nums1
}
