package leetcode

import (
	"github.com/luofeng1/leetcode/link"
)

/**
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

示例：

给定一个链表: 1->2->3->4->5, 和 n = 2.

当删除了倒数第二个节点后，链表变为 1->2->3->5.
说明：

给定的 n 保证是有效的。

进阶：

你能尝试使用一趟扫描实现吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
一次遍历: 需要使用快慢针 来做定位;
	当快针到达结果的时候;
	慢针正好在 剔除 节点的前节点
注意: 使用哨兵节点防止出现删除首节点的问题
*/
func removeNthFromEnd(head *link.ListNode, n int) *link.ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	// 增加虚拟头节点:哨兵节点; 如果不用; 当删除的节点为 head 节点时候,就会出现异常
	//     总结: 凡是删除节点的操作,都需要引入哨兵节点
	dummyHead := &link.ListNode{Next: head}

	quickNode := dummyHead
	slowNode := dummyHead

	// 快指针的Next == nil 则说 快指针已经到了 最后一个节点,此时 slowNode 为要删除的节点的前一个节点
	for quickNode.Next != nil {
		quickNode = quickNode.Next
		if n == 0 {
			slowNode = slowNode.Next
		} else {
			n--
		}
	}
	// fmt.Printf("慢指针: %v 快指针: %v\n", slowNode.ToArray(), quickNode.ToArray())
	slowNode.Next = slowNode.Next.Next
	return dummyHead.Next
}
