package leetcode

import (
	"fmt"
	"testing"

	"github.com/luofeng1/leetcode/link"
)

func Test_addTwoNumbersII(t *testing.T) {
	head1 := link.NewList([]int{7, 2, 4, 3})
	head2 := link.NewList([]int{5, 6, 4})
	head3 := link.NewList([]int{7, 8, 0, 7})

	head4 := link.NewList([]int{7, 2, 4, 3})
	head5 := link.NewList([]int{8, 5, 6, 4})
	head6 := link.NewList([]int{1, 5, 8, 0, 7})
	type args struct {
		l1 *link.ListNode
		l2 *link.ListNode
	}
	tests := []struct {
		name string
		args args
		want *link.ListNode
	}{
		{
			name: "test1",
			args: args{
				head1, head2,
			},
			want: head3,
		},
		{
			name: "test2",
			args: args{
				head4, head5,
			},
			want: head6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addTwoNumbersII(tt.args.l1, tt.args.l2)
			compareNode := tt.want
			fmt.Println("got: ", got.ToArray())
			fmt.Println("compareNode: ", tt.want.ToArray())
			for node := compareNode; node != nil; node = node.Next {
				if got == nil || got.Val != node.Val {
					t.Errorf("addTwoNumbersII() = %v, want %v", got, tt.want)
					break
				}
				got = got.Next
			}
		})
	}
}
