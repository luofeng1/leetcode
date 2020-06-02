package leetcode

import (
	"fmt"

	"github.com/luofeng1/leetcode/tree"
)

/**
给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。

示例:

输入: [1,2,3,null,5,null,4]
输出: [1, 3, 4]
解释:

   1            <---
 /   \
2     3         <---
 \     \
  5     4       <---

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-right-side-view
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
func rightSideView(root *tree.TreeNode) []int {
	return rightSideViewByBFS(root) // 广度搜索
	// return rightSideViewByDfS(root) // 深度搜索
}

// rightSideViewByBFS 广度优先搜索
// 层层遍历,每层最后一个元素即结果
func rightSideViewByBFS(root *tree.TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var result []int
	curLevelNodes := []*tree.TreeNode{root}
	for level := 0; len(curLevelNodes) > 0; level++ {
		var nextLevelNodes []*tree.TreeNode
		for j := 0; j < len(curLevelNodes); j++ {
			node := curLevelNodes[j]
			if j == len(curLevelNodes)-1 {
				result = append(result, node.Val)
			}
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

// rightSideViewByDFS 深度优先搜索
func rightSideViewByDFS(root *tree.TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var result []int
	rightSideViewByDFSInLast(root, 0, &result)
	return result
}

// 由于golang 切片是值拷贝, 我们传递切片的指针,保证始终都是一个切片
func rightSideViewByDFSInLast(root *tree.TreeNode, depth int, depthLastVal *[]int) {
	if root == nil {
		return
	}
	// fmt.Printf("val:%d depth:%d \n", root.Val, depth)
	if len(*depthLastVal) == depth {
		*depthLastVal = append(*depthLastVal, root.Val)
		// fmt.Println("insert last val", root.Val, depthLastVal)
	}
	depth++
	rightSideViewByDFSInLast(root.Right, depth, depthLastVal)
	rightSideViewByDFSInLast(root.Left, depth, depthLastVal)
	return
}

// 练习深度搜索打印
func rightSideViewByDFSPrint(root *tree.TreeNode, depth int) {
	if root == nil {
		return
	}
	fmt.Printf("val:%d depth:%d \n", root.Val, depth)
	depth++
	rightSideViewByDFSPrint(root.Right, depth)
	rightSideViewByDFSPrint(root.Left, depth)
	return
}
