package leetcode

import (
	"math"

	"github.com/luofeng1/leetcode/tree"
)

/*
给定一个二叉树，判断其是否是一个有效的二叉搜索树。

假设一个二叉搜索树具有如下特征：

节点的左子树只包含小于当前节点的数。
节点的右子树只包含大于当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。
示例 1:

输入:
    2
   / \
  1   3
输出: true
示例 2:

输入:
    5
   / \
  1   4
     / \
    3   6
输出: false
解释: 输入为: [5,1,4,null,null,3,6]。
     根节点的值为 5 ，但是其右子节点值为 4 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/validate-binary-search-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *tree.TreeNode) bool {
	result := isValidBSTByMiddle(root)
	if len(result) <= 1 {
		return true
	}
	// fmt.Println("中序:", result) // 利用中序遍历的结果 判断
	for i := 1; i < len(result); i++ {
		if result[i] <= result[i-1] {
			return false
		}
	}
	return true
}

// isValidBSTByMiddle 中序遍历
func isValidBSTByMiddle(root *tree.TreeNode) []int {
	if root == nil {
		return []int{}
	}
	return append(append(isValidBSTByMiddle(root.Left), root.Val), isValidBSTByMiddle(root.Right)...)
}

// 利用递归解
func isValidBST_2(root *tree.TreeNode) bool {
	return isValidBSTByRecursion(root, math.MinInt64, math.MaxInt64)
}

// 始终提供最低值&最高值
func isValidBSTByRecursion(root *tree.TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val >= upper || root.Val <= lower {
		return false
	}
	return isValidBSTByRecursion(root.Left, lower, root.Val) && isValidBSTByRecursion(root.Right, root.Val, upper)
}
