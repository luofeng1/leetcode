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
