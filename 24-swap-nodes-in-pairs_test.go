package leetcode

import (
	"fmt"
	"testing"

	"github.com/luofeng1/leetcode/link"
)

func Test_swapPairs(t *testing.T) {
	head1 := link.NewList([]int{1, 2, 3, 4})
	head2 := link.NewList([]int{2, 1, 4, 3})
	head3 := link.NewList([]int{1, 2, 3})
	head4 := link.NewList([]int{2, 1, 3})
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
			args: args{head: head1},
			want: head2,
		},
		{
			name: "test2",
			args: args{head: head3},
			want: head4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := swapPairs(tt.args.head)
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
				t.Errorf("swapPairs() = %v, want %v", got, tt.want)
			}
			for node := compareNode; node != nil; node = node.Next {
				if got == nil || got.Val != node.Val {
					t.Errorf("swapPairs() = %v, want %v", got, tt.want)
					break
				}
				got = got.Next
			}
		})
	}
}
