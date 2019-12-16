package leetcode

import (
	"github.com/luofeng1/leetcode/link"
)

/**
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
通过单链表的定义可以得知，单链表也是递归结构，因此，也可以使用递归的方式来进行reverse操作。

由于单链表是线性的，使用递归方式将导致栈的使用也是线性的，当链表长度达到一定程度时，递归会导致爆栈，因此，现实中并不推荐使用递归方式来操作链表。
*/
func addTwoNumbers(l1 *link.ListNode, l2 *link.ListNode) *link.ListNode {
	result := &link.ListNode{Val: 0}
	current := result
	var carry int // 用于代步进位

	for l1 != nil || l2 != nil {
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		sum += carry

		node := &link.ListNode{Val: sum % 10}
		current.Next = node
		current = node
		carry = sum / 10
	}

	if carry > 0 {
		current.Next = &link.ListNode{Val: carry}
	}
	return result.Next
}
