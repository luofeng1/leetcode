package tree

// TreeNode 树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func (node *TreeNode) InOrder(result *[]int) {
//	if node.Left != nil {
//		node.Left.InOrder(result)
//	}
//	*result = append(*result, node.Val)
//	if node.Right != nil {
//		node.Right.InOrder(result)
//	}
//}

const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX

func NewTree(arr []int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}
	head := &TreeNode{Val: arr[0]}
	parentsNode := []*TreeNode{head}
	currentsNode := []*TreeNode{head}
	depth := 1
	for i := 1; i < len(arr); i++ {
		if isNewDepth(i, depth) {
			depth += 1
			parentsNode = currentsNode
			currentsNode = []*TreeNode{}
		}
		currentParent := getParentNode(parentsNode, depth, i)
		if currentParent == nil {
			continue
		}
		currentNode := getNodeVal(arr[i])
		if i%2 != 0 {
			currentParent.Left = currentNode // 处理左
		} else {
			currentParent.Right = currentNode // 处理右
		}
		currentsNode = append(currentsNode, currentNode)
	}
	return head
}

func isNewDepth(index, depth int) bool {
	return index+1 == 2<<(depth-1)
}

func getParentNode(parentsNode []*TreeNode, depth, index int) *TreeNode {
	preIndex := 2<<(depth-2) - 1
	currentIndex := (index - preIndex) / 2
	return parentsNode[currentIndex]
}

func getNodeVal(v int) *TreeNode {
	if v == INT_MIN {
		return nil
	}
	return &TreeNode{Val: v}
}
