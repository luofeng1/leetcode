package leetcode

import (
	"github.com/luofeng1/leetcode/link"
)

/*
反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。

说明:
1 ≤ m ≤ n ≤ 链表长度。

示例:

输入: 1->2->3->4->5->NULL, m = 2, n = 4
输出: 1->4->3->2->5->NULL

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-linked-list-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func reverseBetween(head *link.ListNode, m int, n int) *link.ListNode {
	/**
	输入: 0->1->2->3->4->5->NULL, m = 3, n = 5
	输出: 0->1->4->3->2->5->NULL
	*/
	if head == nil || head.Next == nil {
		return head
	}
	if m == 1 {
		// 返回新的 头节点; 处理范围为 [2,3,4,5] 返回的为[4,3,2,5]的头结点4
		return reverseBetweenBefore(head, n)
	}

	// 入参是 1.Next 即 2;
	// m==1开始返回的第一个是: [4,3,2,5] 的 4 节点; 作为 1.Next 的节点 即: [1,4,3,2,5]
	// 重复返回 拼接剩余的head [0,1,4,3,2,5]
	head.Next = reverseBetween(head.Next, m-1, n-1)
	return head
}

func reverseBetweenBefore(head *link.ListNode, n int) *link.ListNode {
	/**
	输入: 2->3->4->5->NULL, n = 3
	输出: 4->3->2->5->NULL
	*/
	if n == 1 {
		// 当 n==1 返回的是4节点
		return head
	}
	last := reverseBetweenBefore(head.Next, n-1)
	// 当 n==1 返回的是4节点; 此时 head 是 3 节点;
	// 需要把 4.Next 节点替换成 3 => head.Next.Next = head
	// 同时 需要 把 3.Next 节点 挂在 5之后 => head.Next = head.Next.Next (备注此步骤需要和上面一起进行;否则需要重新赋值)
	head.Next.Next, head.Next = head, head.Next.Next
	return last
}

func reverseBetween2(head *link.ListNode, m int, n int) *link.ListNode {
	/**
	输入: 0->1->2->3->4->5->NULL, m = 3, n = 5
	输出: 0->1->4->3->2->5->NULL
	*/
	if head == nil || head.Next == nil || m == n {
		return head
	}
	cur := head             // m 节点; 始终为当前节点
	var tail *link.ListNode // 尾结点 m 节点的前一个节点
	for m != 1 {
		tail, cur = cur, cur.Next // m的前一个节点 即 tail; cur 即m节点
		m--                       // 减掉前面的节点
		n--                       // 随着m继续减
	}

	// m开始后面的翻转后的头结点
	if afterHead := reverseBetweenBefore2(cur, n); tail == nil {
		// tail == nil 说明是从第一个节点开始翻转的;也就是说: head 应该为新的head
		return afterHead
	} else {
		// 翻转后的头结点 即 m的前一个节点 tail 的下一个节点
		tail.Next = afterHead
	}

	return head
}

func reverseBetweenBefore2(head *link.ListNode, n int) *link.ListNode {
	/**
	输入: 2->3->4->5->NULL, n = 3
	输出: 4->3->2->5->NULL
	*/
	//pre := &link.ListNode{Next: head}
	cur := head
	var prev *link.ListNode
	for n != 0 {
		cur, prev, cur.Next = cur.Next, cur, prev
		// 当为最后一个节点的时候; 参数 head 节点 即为 最后一个节点;
		//     此时,将剩余的节点也就是 cur 节点拼接在 head.Next 上面; 即完成了翻转+拼接
		if n == 1 {
			head.Next = cur
		}
		n--
	}
	return prev
}
