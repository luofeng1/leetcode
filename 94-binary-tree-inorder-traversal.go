package leetcode

import (
	"github.com/luofeng1/leetcode/stack"
	"github.com/luofeng1/leetcode/tree"
)

/**
给定一个二叉树，返回它的中序 遍历。

示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,3,2]
进阶: 递归算法很简单，你可以通过迭代算法完成吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-inorder-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 中序遍历 [ [左子树], 根节点, [右子树] ]

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 迭代实现遍历 使用栈
func inorderTraversal(root *tree.TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var result []int
	curNode := root
	stackNodes := stack.NewStack()
	stackNodes.Push(root)
	for stackNodes.Len() != 0 {
		if curNode.Left != nil {
			stackNodes.Push(curNode.Left)
			curNode = curNode.Left
			continue
		}

		node := stackNodes.Pop().(*tree.TreeNode)
		result = append(result, node.Val)

		if node.Right != nil {
			stackNodes.Push(node.Right)
			curNode = node.Right
		}
	}
	return result
}

// 迭代实现遍历 使用栈 & 颜色标记 NB
/*使用颜色标记节点的状态，新节点为白色，已访问的节点为灰色。
如果遇到的节点为白色，则将其标记为灰色，然后将其右子节点、自身、左子节点依次入栈。
如果遇到的节点为灰色，则将节点的值输出。

作者：hzhu212
链接：https://leetcode-cn.com/problems/binary-tree-inorder-traversal/solution/yan-se-biao-ji-fa-yi-chong-tong-yong-qie-jian-ming/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
// 可以通过判断数据类型替换颜色; 不够语义
func inorderTraversalByColorInStack(root *tree.TreeNode) []int {
	if root == nil {
		return []int{}
	}
	const (
		write int = iota
		black
	)
	type color struct {
		node  *tree.TreeNode
		color int
	}
	var result []int
	stackNodes := stack.NewStack()
	stackNodes.Push(&color{root, write})

	for stackNodes.Len() != 0 {
		node := stackNodes.Pop().(*color)
		// fmt.Println("node: ", node.node)
		if node.node == nil {
			continue
		}
		if node.color == write {
			stackNodes.Push(&color{node.node.Right, write})
			stackNodes.Push(&color{node.node, black})
			stackNodes.Push(&color{node.node.Left, write})
		} else {
			result = append(result, node.node.Val)
		}
	}
	return result
}

// 递归实现遍历
func inorderTraversalByRecursion(root *tree.TreeNode) []int {
	if root == nil {
		return []int{}
	}
	return append(append(inorderTraversal(root.Left), root.Val), inorderTraversal(root.Right)...)
}
