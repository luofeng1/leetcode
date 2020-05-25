package leetcode

import (
	"fmt"
	"testing"

	"github.com/luofeng1/leetcode/link"
)

func Test_sortList(t *testing.T) {
	head1 := link.NewList([]int{5, 2, 3, 7, 9, 1, 8})
	head2 := link.NewList([]int{1, 2, 3, 5, 7, 8, 9})

	head3 := link.NewList([]int{-1, 5, 3, 4, 0})
	head4 := link.NewList([]int{-1, 0, 3, 4, 5})
	type args struct {
		head *link.ListNode
	}
	tests := []struct {
		name string
		args args
		want *link.ListNode
	}{
		{
			name: "排序",
			args: args{head1},
			want: head2,
		},
		{
			name: "排序2",
			args: args{head3},
			want: head4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sortList(tt.args.head)
			compareNode := tt.want
			for node := compareNode; node != nil; node = node.Next {
				if got == nil || got.Val != node.Val {
					t.Errorf("sortList() = %v, want %v", got, tt.want)
					break
				}
				got = got.Next
			}
		})
	}
}

func Test_sortListCut(t *testing.T) {
	head1 := link.NewList([]int{1, 7, 2, 3})
	head2 := link.NewList([]int{2, 3})
	type args struct {
		head   *link.ListNode
		length int
	}
	tests := []struct {
		name string
		args args
		want *link.ListNode
	}{
		{
			name: "cut",
			args: args{head1, 2},
			want: head2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println("all: ", tt.args.head.ToArray())
			got := sortListCut(tt.args.head, tt.args.length)
			fmt.Println("result: ", got.ToArray())
			fmt.Println("cut: ", tt.args.head.ToArray())
			compareNode := tt.want
			for node := compareNode; node != nil; node = node.Next {
				if got == nil || got.Val != node.Val {
					t.Errorf("sortListCut() = %v, want %v", got, tt.want)
					break
				}
				got = got.Next
			}
		})
	}
}

func Test_sortListMerge(t *testing.T) {
	head1 := link.NewList([]int{1, 7, 9})
	head2 := link.NewList([]int{2})
	head3 := link.NewList([]int{1, 2, 7, 9})
	type args struct {
		head1 *link.ListNode
		head2 *link.ListNode
	}
	tests := []struct {
		name string
		args args
		want *link.ListNode
	}{
		{
			name: "cut",
			args: args{head1, head2},
			want: head3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := sortListMerge(tt.args.head1, tt.args.head2)
			compareNode := tt.want
			for node := compareNode; node != nil; node = node.Next {
				if got == nil || got.Val != node.Val {
					t.Errorf("sortListCut() = %v, want %v", got, tt.want)
					break
				}
				got = got.Next
			}
		})
	}
}
