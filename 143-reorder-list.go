package leetcode

import (
	"github.com/luofeng1/leetcode/link"
)

/**
给定一个单链表 L：L0→L1→…→Ln-1→Ln ，
将其重新排列后变为： L0→Ln→L1→Ln-1→L2→Ln-2→…

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

示例 1:

给定链表 1->2->3->4, 重新排列为 1->4->2->3.
示例 2:

给定链表 1->2->3->4->5, 重新排列为 1->5->2->4->3.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reorder-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func reorderList(head *link.ListNode) *link.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	listLength := 0
	for node := head; node != nil; node = node.Next {
		listLength++
	}
	head1, head2 := reorderListCut(head, listLength/2)
	return reorderListMerge(head1, reorderListReverse(head2))
}

// reorderListCut 切割链表 -> [1,2,3,4,5] --> [1,2] [3,4,5]
func reorderListCut(head *link.ListNode, position int) (*link.ListNode, *link.ListNode) {
	head2 := head
	for position != 0 {
		if position == 1 {
			head2, head2.Next = head2.Next, nil
		} else {
			head2 = head2.Next
		}
		position--
	}
	// fmt.Println(head.ToArray(), head2.ToArray())
	return head, head2
}

// reorderListReverse 反转链表 [3,4,5] -> [5,4,3]
func reorderListReverse(head *link.ListNode) *link.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur := head
	var pre *link.ListNode
	for cur != nil {
		cur, cur.Next, pre = cur.Next, pre, cur
	}
	return pre
}

// reorderListMerge 合并链表直接合并交叉排序 [1,2] [5,4,3] -> [1,5,2,4,3]
func reorderListMerge(head1 *link.ListNode, head2 *link.ListNode) *link.ListNode {
	dummyNode := &link.ListNode{Next: head1}
	for head1.Next != nil {
		head1.Next, head2.Next, head1, head2 = head2, head1.Next, head1.Next, head2.Next
	}
	head1.Next = head2
	return dummyNode.Next
}
