package leetcode

import (
	"fmt"
	"testing"

	"github.com/luofeng1/leetcode/link"
)

func Test_reorderList(t *testing.T) {
	head1 := link.NewList([]int{1, 2, 3, 4, 5})
	head2 := link.NewList([]int{1, 5, 2, 4, 3})
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reorderList(tt.args.head)
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
				t.Errorf("reorderList() = %v, want %v", got, tt.want)
			}
			for node := compareNode; node != nil; node = node.Next {
				if got == nil || got.Val != node.Val {
					t.Errorf("reorderList() = %v, want %v", got, tt.want)
					break
				}
				got = got.Next
			}
		})
	}
}

func Test_reorderListCut(t *testing.T) {
	head1 := link.NewList([]int{1, 2, 3, 4, 5})
	head2 := link.NewList([]int{1, 2})
	head3 := link.NewList([]int{3, 4, 5})
	type args struct {
		head     *link.ListNode
		position int
	}
	tests := []struct {
		name  string
		args  args
		want  *link.ListNode
		want1 *link.ListNode
	}{
		{
			name:  "test1",
			args:  args{head: head1, position: 2},
			want:  head2,
			want1: head3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := reorderListCut(tt.args.head, tt.args.position)
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
				t.Errorf("reorderListCut() got = %v, want %v", got, tt.want)
			}
			for node := compareNode; node != nil; node = node.Next {
				if got == nil || got.Val != node.Val {
					t.Errorf("reorderListCut() = %v, want %v", got, tt.want)
					break
				}
				got = got.Next
			}

			compareNode1 := tt.want1
			fmt.Println("got1: ", got1.ToArray())
			fmt.Println("compareNode1: ", tt.want1.ToArray())
			gotLength1 := 0
			compareNodeLength1 := 0
			for node := got1; node != nil; node = node.Next {
				gotLength1++
			}
			for node := compareNode1; node != nil; node = node.Next {
				compareNodeLength1++
			}
			if gotLength1 != compareNodeLength1 {
				t.Errorf("reorderListCut1() = %v, want %v", got1, tt.want1)
			}
			for node := compareNode1; node != nil; node = node.Next {
				if got1 == nil || got1.Val != node.Val {
					t.Errorf("reorderListCut1() got1 = %v, want %v", got1, tt.want1)
					break
				}
				got1 = got1.Next
			}
		})
	}
}

func Test_reorderListReverse(t *testing.T) {
	head1 := link.NewList([]int{3, 2, 1})
	head2 := link.NewList([]int{1, 2, 3})
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reorderListReverse(tt.args.head)
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
				t.Errorf("reorderListReverse() = %v, want %v", got, tt.want)
			}
			for node := compareNode; node != nil; node = node.Next {
				if got == nil || got.Val != node.Val {
					t.Errorf("reorderListReverse() = %v, want %v", got, tt.want)
					break
				}
				got = got.Next
			}
		})
	}
}

func Test_reorderListMerge(t *testing.T) {
	head1 := link.NewList([]int{1, 2})
	head2 := link.NewList([]int{5, 4, 3})
	head3 := link.NewList([]int{1, 5, 2, 4, 3})
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
			name: "test1",
			args: args{head1, head2},
			want: head3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reorderListMerge(tt.args.head1, tt.args.head2)
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
				t.Errorf("reorderListMerge() = %v, want %v", got, tt.want)
			}
			for node := compareNode; node != nil; node = node.Next {
				if got == nil || got.Val != node.Val {
					t.Errorf("reorderListMerge() = %v, want %v", got, tt.want)
					break
				}
				got = got.Next
			}
		})
	}
}
