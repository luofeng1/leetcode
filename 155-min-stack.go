package leetcode

import (
	"container/list"
	"fmt"
	"math"
	"sync"
)

/**
设计一个支持 push，pop，top 操作，并能在常数时间内检索到最小元素的栈。

push(x) -- 将元素 x 推入栈中。
pop() -- 删除栈顶的元素。
top() -- 获取栈顶元素。
getMin() -- 检索栈中的最小元素。
示例:

MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.getMin();   --> 返回 -2.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/min-stack
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
最小栈存储的不应该是真实值，而是真实值和min的差值
top的时候涉及到对数据的还原，这里千万注意是上一个最小值
*/
type MinStack struct {
	min  int
	list *list.List
	lock *sync.RWMutex
}

/** initialize your data structure here. */
func NewMinStack() *MinStack {
	return &MinStack{math.MaxInt32, list.New(), &sync.RWMutex{}}
}

func (stack *MinStack) Push(x int) {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	min := stack.min
	if stack.min > x {
		stack.min = x
	}
	stack.list.PushBack(x - min)
}

func (stack *MinStack) Pop() {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	e := stack.list.Back()
	if e.Value.(int) < 0 {
		stack.min = stack.min - e.Value.(int)
	}
	if e != nil {
		stack.list.Remove(e)
		return
	}
	return
}

func (stack *MinStack) Top() int {
	if stack.list.Back().Value.(int) < 0 {
		fmt.Println("top:", stack.list.Back().Value.(int), stack.min)
		return stack.min
	}
	fmt.Println("top:", stack.list.Back().Value.(int), stack.min)
	return stack.list.Back().Value.(int) + stack.min
}

func (stack *MinStack) GetMin() int {
	return stack.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
