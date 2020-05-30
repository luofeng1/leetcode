package leetcode

import (
	"github.com/luofeng1/leetcode/link"
)

/**
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

示例:

给定 1->2->3->4, 你应该返回 2->1->4->3.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/swap-nodes-in-pairs
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *link.ListNode) *link.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &link.ListNode{Next: head}
	pre := dummy
	cur := head
	/**
	2 -> 1 -> 3 -> 4 -> nil
	pre = 1 cur = 3
			┌--------→
	2 -> 1  3 <- 4  nil
		 └-------→

	pre = 3 cur = nil
	*/
	for cur != nil && cur.Next != nil {
		pre, cur, pre.Next, cur.Next, cur.Next.Next = cur, cur.Next.Next, cur.Next, cur.Next.Next, cur
		// fmt.Println(pre.ToArray(), cur.ToArray(), dummy.ToArray())
	}
	return dummy.Next
}
