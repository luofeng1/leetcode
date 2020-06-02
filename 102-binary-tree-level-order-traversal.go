package leetcode

import (
	"fmt"

	"github.com/luofeng1/leetcode/tree"
)

/**
给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。
示例：
二叉树：[3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：
[
  [3],
  [9,20],
  [15,7]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-level-order-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
「层序遍历」、「最短路径」:广度优先搜索
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *tree.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var result [][]int
	// 当前层的所有结点
	curLevelNodes := []*tree.TreeNode{root}
	for level := 0; len(curLevelNodes) > 0; level++ {
		// 下一层的所有结点
		var nextLevelNode []*tree.TreeNode
		result = append(result, []int{})
		// 遍历当前层的全部结点; 以获取下一层的节点
		for i := 0; i < len(curLevelNodes); i++ {
			node := curLevelNodes[i]
			// fmt.Printf("第%d层 值:%d \n", level, node.Val)
			// 记录值
			result[level] = append(result[level], node.Val)

			// 左右节点存在,则放到下一层
			if node.Left != nil {
				nextLevelNode = append(nextLevelNode, node.Left)
			}
			if node.Right != nil {
				nextLevelNode = append(nextLevelNode, node.Right)
			}
		}
		curLevelNodes = nextLevelNode
	}
	return result
}

// levelOrderInOneLevelByBFS 使用广度优先搜索遍历
func levelOrderInOneLevelByBFS(root *tree.TreeNode) []int {
	var result []int
	nodes := []*tree.TreeNode{root}

	for item := 0; len(nodes) > item; item++ {
		node := nodes[item]
		if node.Val != tree.INT_MIN {
			result = append(result, node.Val)
		}
		if node.Left != nil {
			nodes = append(nodes, node.Left)
		}
		if node.Right != nil {
			nodes = append(nodes, node.Right)
		}
	}
	fmt.Println(result)
	return result
}
