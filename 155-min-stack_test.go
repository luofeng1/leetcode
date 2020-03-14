package leetcode

import (
	"container/list"
	"fmt"
	"sync"
	"testing"
)

func TestMinStack_GetMin(t *testing.T) {
	type fields struct {
		min  int
		list *list.List
		lock *sync.RWMutex
	}
	stack1 := NewMinStack()
	stack1.Push(-2)
	stack1.Push(0)
	stack1.Push(-3)
	stack1.Pop()
	fmt.Println(stack1.Top())

	stack2 := NewMinStack()
	stack2.Push(-1)
	fmt.Println(stack2.Top())
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "getMin",
			fields: fields{
				min:  stack1.min,
				list: stack1.list,
				lock: stack1.lock,
			},
			want: -2,
		},
		{
			name: "getMin2",
			fields: fields{
				min:  stack2.min,
				list: stack2.list,
				lock: stack2.lock,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &MinStack{
				min:  tt.fields.min,
				list: tt.fields.list,
				lock: tt.fields.lock,
			}
			if got := this.GetMin(); got != tt.want {
				t.Errorf("GetMin() = %v, want %v", got, tt.want)
			}
		})
	}
}
