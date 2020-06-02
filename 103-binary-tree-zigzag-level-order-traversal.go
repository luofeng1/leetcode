package leetcode

import (
	"fmt"

	"github.com/luofeng1/leetcode/tree"
)

/**
给定一个二叉树，返回其节点值的锯齿形层次遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

例如：
给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回锯齿形层次遍历如下：

[
  [3],
  [20,9],
  [15,7]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal
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
func zigzagLevelOrder(root *tree.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	curLevelNodes := []*tree.TreeNode{root}
	var result [][]int
	for level := 0; len(curLevelNodes) > 0; level++ {
		var nextLevelNodes []*tree.TreeNode
		result = append(result, []int{})
		for j := 0; j < len(curLevelNodes); j++ {
			node := curLevelNodes[j]
			if level%2 == 0 {
				result[level] = append(result[level], node.Val)
			} else {
				result[level] = append(append([]int{}, node.Val), result[level]...)
			}
			fmt.Printf("层级: %d 值:%d \n", level, node.Val)
			if node.Left != nil {
				nextLevelNodes = append(nextLevelNodes, node.Left)
			}
			if node.Right != nil {
				nextLevelNodes = append(nextLevelNodes, node.Right)
			}
		}
		curLevelNodes = nextLevelNodes
	}
	return result
}
