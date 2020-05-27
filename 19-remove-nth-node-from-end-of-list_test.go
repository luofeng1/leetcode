package leetcode

import (
	"fmt"
	"testing"

	"github.com/luofeng1/leetcode/link"
)

func Test_removeNthFromEnd(t *testing.T) {
	head1 := link.NewList([]int{1, 2, 3, 4, 5})
	head2 := link.NewList([]int{1, 2, 3, 5})
	head3 := link.NewList([]int{1, 2})
	head4 := link.NewList([]int{2})
	type args struct {
		head *link.ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *link.ListNode
	}{
		{
			name: "test1",
			args: args{
				head: head1,
				n:    2,
			},
			want: head2,
		},
		{
			name: "头结点异常:test2",
			args: args{
				head: head3,
				n:    2,
			},
			want: head4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeNthFromEnd(tt.args.head, tt.args.n)
			compareNode := tt.want
			fmt.Println("got: ", got.ToArray())
			fmt.Println("compareNode: ", tt.want.ToArray())
			for node := compareNode; node != nil; node = node.Next {
				if got == nil || got.Val != node.Val {
					t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
					break
				}
				got = got.Next
			}
		})
	}
}
