package leetcode

import (
	"fmt"
	"testing"

	"github.com/luofeng1/leetcode/link"
)

func Test_deleteDuplicates(t *testing.T) {
	head1 := link.NewList([]int{1, 2, 3, 3, 4, 4, 5})
	head2 := link.NewList([]int{1, 2, 5})
	head3 := link.NewList([]int{1, 1, 1, 2, 3})
	head4 := link.NewList([]int{2, 3})
	head5 := link.NewList([]int{2})
	head6 := link.NewList([]int{2, 2})
	head7 := link.NewList([]int{})
	type args struct {
		head *link.ListNode
	}
	tests := []struct {
		name string
		args args
		want *link.ListNode
	}{
		{
			name: "test1",
			args: args{head1},
			want: head2,
		},
		{
			name: "test2",
			args: args{head3},
			want: head4,
		},
		{
			name: "test3",
			args: args{head5},
			want: head5,
		},
		{
			name: "test4",
			args: args{head6},
			want: head7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := deleteDuplicates(tt.args.head)
			compareNode := tt.want
			fmt.Println("got: ", got.ToArray())
			fmt.Println("compareNode: ", tt.want.ToArray())
			gotLength := 0
			compareNodeLength := 0
			for node := got; node != nil; node = node.Next {
				gotLength++
			}
			for node := compareNode; node != nil; node = node.Next {
				compareNodeLength++
			}
			if gotLength != compareNodeLength {
				t.Errorf("deleteDuplicates() = %v, want %v", got, tt.want)
			}
			for node := compareNode; node != nil; node = node.Next {
				if got == nil || got.Val != node.Val {
					t.Errorf("deleteDuplicates() = %v, want %v", got, tt.want)
					break
				}
				got = got.Next
			}
		})
	}
}
