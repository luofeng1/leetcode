package leetcode

import "github.com/luofeng1/leetcode/link"

/**
反转一个单链表。

示例:

输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL
进阶:
你可以迭代或递归地反转链表。你能否用两种方法解决这道题？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/**
2 为例:
因为: 2.Next = 1
cur = 2
pre = 1
cur.Next = 1 : pre
pre = 2 cur
cur = 3 cur.Next 需提前保存
如果 cur = nil 则停止;

链表思路:从中间获取一个;看如何去操作
*/
//func reverseList(head *link.ListNode) *link.ListNode {
//	cur := head
//	var prev *link.ListNode
//
//	for cur != nil {
//		c
//	}
//	return prev
//}

/**
输入: [1,2,3,4,5]
输出: [5,4,3,2,1]
预期结果: [5,4,3,2,1]
*/

// 递归解题
func reverseList(head *link.ListNode) *link.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	last := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}
