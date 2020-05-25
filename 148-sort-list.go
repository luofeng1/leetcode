package leetcode

import (
	"fmt"

	"github.com/luofeng1/leetcode/link"
)

/**
在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序。

示例 1:

输入: 4->2->1->3
输出: 1->2->3->4
示例 2:

输入: -1->5->3->4->0
输出: -1->0->3->4->5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sort-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *link.ListNode) *link.ListNode {
	// 1. 排除头部为开 和 只有1个节点的情况
	if head == nil || head.Next == nil {
		return head
	}

	// 2. 获取链表长度
	listLength := 0
	for node := head; node != nil; node = node.Next {
		listLength++
	}
	//fmt.Printf("链表长度: %d \n", listLength)

	// 3. 引入新头节点 & 层级标识
	newHead := &link.ListNode{Next: head}
	//fmt.Printf("第%d层 result: %v \n", 0, newHead.Next.ToArray())
	for storey := 1; storey <= listLength; storey *= 2 {
		cur := newHead.Next // 当前结点
		tail := newHead     // 尾结点
		for cur != nil {
			left := cur
			right := sortListCut(left, storey)
			cur = sortListCut(right, storey)
			tail.Next, tail = sortListMerge(left, right) // 拼接后返回的头结点正好是当前尾结点的next结点
		}
		//fmt.Printf("第%d层 result: %v \n", storey, newHead.Next.ToArray())
	}

	return newHead.Next
}

// sortListMerge 合并2个链表 返回合并后的头部结点
func sortListMerge(head1 *link.ListNode, head2 *link.ListNode) (*link.ListNode, *link.ListNode) {
	head := &link.ListNode{}
	pre := head
	for head1 != nil && head2 != nil {
		if head1.Val > head2.Val {
			pre.Next = head2
			pre = pre.Next
			head2 = head2.Next
		} else {
			pre.Next = head1
			pre = pre.Next
			head1 = head1.Next
		}
	}
	tail := pre
	if head1 == nil {
		pre.Next = head2
	} else {
		pre.Next = head1
	}
	for tail.Next != nil {
		tail = tail.Next
	}
	fmt.Println("head: ", head.ToArray(), tail.ToArray())
	return head.Next, tail
}

// sortListCut 断链操作: 切断某个链表 返回下一段的头结点
func sortListCut(head *link.ListNode, length int) *link.ListNode {
	if head == nil {
		return nil
	}
	var nextPre *link.ListNode
	for length > 1 {
		if head = head.Next; head == nil {
			return nil
		}
		length--
	}
	// pre 新增伪头结点; head.Next = nil 截断链表
	nextPre, head.Next = head.Next, nil
	return nextPre
}

//func sortList1(head *link.ListNode) *link.ListNode {
//	if head == nil || head.Next == nil {
//		return head
//	}
//
//	length := 0
//	for node := head; node != nil; node = node.Next {
//		length++
//	}
//
//	newHead := &link.ListNode{Next: head}
//	for storey := 1; storey < length; storey *= 2 {
//		cur := newHead.Next
//		tail := newHead
//		for cur != nil {
//			left := cur
//			right := sortListCut(left, storey)
//			cur =  sortListCut(right, storey)
//			tail.Next = sortListMerge(left, right)
//			for tail.Next != nil {
//				tail = tail.Next
//			}
//		}
//	}
//
//	return newHead.Next
//}
