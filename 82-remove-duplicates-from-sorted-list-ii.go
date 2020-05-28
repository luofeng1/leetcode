package leetcode

import (
	"github.com/luofeng1/leetcode/link"
)

/**
给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。

示例 1:

输入: 1->2->3->3->4->4->5
输出: 1->2->5
示例 2:

输入: 1->1->1->2->3
输出: 2->3

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
示例1: 运行打印:
pre:  [1 2 3 3 4 4 5] [0 1 2 3 3 4 4 5]
移位:
end:  [2 3 3 4 4 5] [1 2 3 3 4 4 5]
pre:  [2 3 3 4 4 5] [1 2 3 3 4 4 5]
移位:
end:  [3 3 4 4 5] [2 3 3 4 4 5]
pre:  [3 3 4 4 5] [2 3 3 4 4 5]
大跳:
end:  [4 4 5] [2 4 4 5]
pre:  [4 4 5] [2 4 4 5]
大跳:
end:  [5] [2 5]
 */
func deleteDuplicates(head *link.ListNode) *link.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 可能会删除头节点则使用虚拟头结点 哨兵节点
	dummyHead := &link.ListNode{Next: head} // 虚拟头节点
	curNode := dummyHead.Next               // 当前节点
	slowNode := dummyHead                   // 不重复的最后一个节点, 即慢节点

	for curNode != nil && curNode.Next != nil {
		//fmt.Println("pre: ", curNode.ToArray(), slowNode.ToArray())
		curValue := curNode.Val // 当前节点的值
		// 当前节点的值不等于下一个节点的值: 则进行 移位
		//     如果相等 则进行 for 循环到下一个不等的值 大跳
		if curValue != curNode.Next.Val {
			//fmt.Println("移位: ")
			curNode = curNode.Next
			slowNode = slowNode.Next
		} else {
			//fmt.Println("大跳: ")
			for node := curNode; node != nil && node.Val == curValue; node = node.Next {
				curNode = curNode.Next
			}
			// 将新出现的不等值 拼接到slowNode后边
			slowNode.Next = curNode
		}
		//fmt.Println("end: ", curNode.ToArray(), slowNode.ToArray())
	}

	return dummyHead.Next
}
