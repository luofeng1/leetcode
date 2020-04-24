package link

import (
	"reflect"
	"testing"
)

func TestNewList(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "2->4->3",
			args: args{
				arr: []int{2, 4, 3},
			},
			want: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
		},
		{
			name: "nil",
			args: args{
				arr: []int{},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewList(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListNode_ToArray(t *testing.T) {
	head := NewList([]int{1, 2, 3, 4, 5})
	tests := []struct {
		name string
		want []int
	}{
		{
			name: "test1",
			want: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := head
			if got := l.ToArray(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
