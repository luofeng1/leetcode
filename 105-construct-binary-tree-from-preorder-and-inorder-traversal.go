package leetcode

import "github.com/luofeng1/leetcode/tree"

/**
根据一棵树的前序遍历与中序遍历构造二叉树。

注意:
你可以假设树中没有重复的元素。

例如，给出

前序遍历 preorder = [3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]
返回如下的二叉树：

    3
   / \
  9  20
    /  \
   15   7

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/**
前序: [ 根节点, [左子树的前序遍历结果], [右子树的前序遍历结果] ]
中序: [ [左子树的中序遍历结果], 根节点, [右子树的中序遍历结果] ]
获取的目标是
	前序左子树 & 中序左子树
	前序右子树 & 中序右子树
 */
func buildTree(preorder []int, inorder []int) *tree.TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &tree.TreeNode{Val: preorder[0]}
	rootPosition := 0
	for ; rootPosition < len(inorder); rootPosition++ {
		if inorder[rootPosition] == root.Val {
			break
		}
	}
	root.Left = buildTree(preorder[1:rootPosition+1], inorder[:rootPosition])
	root.Right = buildTree(preorder[rootPosition+1:], inorder[rootPosition+1:])
	return root
}
