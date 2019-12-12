package array

/**
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

说明：

你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？

示例 1:

输入: [2,2,1]
输出: 1
示例 2:

输入: [4,1,2,1,2]
输出: 4
*/

// hash 运算
func singleNumber1(nums []int) int {
	o := map[int]int{}
	for _, v := range nums {
		if _, exists := o[v]; exists {
			delete(o, v)
		} else {
			o[v] = 1
		}
	}
	for k := range o {
		return k
	}
	return 0
}

// 位运算
/**
方法 4：位操作
概念

如果我们对 0 和二进制位做 XOR 运算，得到的仍然是这个二进制位
a ⊕ 0 = aa⊕0=a
如果我们对相同的二进制位做 XOR 运算，返回的结果是 0
a ⊕ a = 0a⊕a=0
XOR 满足交换律和结合律
a ⊕ b ⊕ a = (a ⊕ a) ⊕ b = 0 ⊕ b = b = a⊕b⊕a=(a⊕a)⊕b=0⊕b=b
所以我们只需要将所有的数进行 XOR 操作，得到那个唯一的数字。

Python
class Solution(object):
    def singleNumber(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        a = 0
        for i in nums:
            a ^= i
        return a
复杂度分析

时间复杂度： O(n)O(n) 。我们只需要将 \text{nums}nums 中的元素遍历一遍，所以时间复杂度就是 \text{nums}nums 中的元素个数。
空间复杂度：O(1)O(1) 。

5的二进制位是0000 0101 ， 9的二进制位是0000 1001，也就是0101 ^ 1001,结果为1100 , 00001100的十进制位是12
*/
func singleNumber(nums []int) int {
	var o int
	for _, v := range nums {
		o = v ^ o
	}
	return o
}
