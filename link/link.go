package link

type ListNode struct {
	Val  int
	Next *ListNode
}

// NewList 实现一个list
func NewList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	head := &ListNode{Val: arr[0]}
	parent := head
	for i := 1; i < len(arr); i++ {
		parent.Next = &ListNode{Val: arr[i]}
		parent = parent.Next
	}
	return head
}

func (l *ListNode) ToArray() []int {
	var arr []int
	cur := l
	for cur != nil {
		arr = append(arr, cur.Val)
		cur = cur.Next
	}
	return arr
}
