package leetcode

import (
	"github.com/luofeng1/leetcode/link"
)

/**
给你两个 非空 链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。

你可以假设除了数字 0 之外，这两个数字都不会以零开头。

进阶：

如果输入链表不能修改该如何处理？换句话说，你不能对列表中的节点进行翻转。

示例：

输入：(7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 8 -> 0 -> 7

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 双栈处理
func addTwoNumbersII(l1 *link.ListNode, l2 *link.ListNode) *link.ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var stack1 []int
	var stack2 []int
	for l1 != nil {
		stack1 = append(stack1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		stack2 = append(stack2, l2.Val)
		l2 = l2.Next
	}
	//fmt.Println("stack1: ", stack1)
	//fmt.Println("stack2: ", stack2)

	head := &link.ListNode{} // 虚拟头结点
	carry := 0               // 进位使用
	position := 0            // 遍历到了第几位
	// 注意: carry > 0; 需要继续增加;
	// 判断只要某个还有值则继续查询
	for position < len(stack1) || position < len(stack2) || carry > 0 {
		sum := carry
		if position < len(stack1) {
			sum += stack1[len(stack1)-position-1]
		}
		if position < len(stack2) {
			sum += stack2[len(stack2)-position-1]
		}
		carry = sum / 10
		// fmt.Printf("进位值: %d 当前位值: %d\n", carry, sum%10)
		newNode := &link.ListNode{Val: sum % 10}
		head.Next, newNode.Next = newNode, head.Next
		position++
	}
	return head.Next
}
