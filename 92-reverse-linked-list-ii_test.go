package leetcode

import (
	"fmt"
	"testing"

	"github.com/luofeng1/leetcode/link"
)

func Test_reverseBetween(t *testing.T) {
	head1 := link.NewList([]int{1, 2, 3, 4, 5})
	head2 := link.NewList([]int{1, 4, 3, 2, 5})
	//head2 := link.NewList([]int{4, 3, 2, 1, 5})
	type args struct {
		head *link.ListNode
		m    int
		n    int
	}
	tests := []struct {
		name string
		args args
		want *link.ListNode
	}{
		{
			name: "test1",
			args: args{head: head1, m: 2, n: 4},
			want: head2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseBetween(tt.args.head, tt.args.m, tt.args.n)
			compareNode := tt.want
			fmt.Println("got: ", got.ToArray())
			fmt.Println("compareNode: ", tt.want.ToArray())
			for node := compareNode; node != nil; node = node.Next {
				if got == nil || got.Val != node.Val {
					t.Errorf("reverseBetween() = %v, want %v", got, tt.want)
					break
				}
				got = got.Next
			}
		})
	}
}

func Test_reverseBetweenBefore(t *testing.T) {
	head1 := link.NewList([]int{1, 2, 3, 4, 5})
	head2 := link.NewList([]int{4, 3, 2, 1, 5})
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
			args: args{head: head1, n: 4},
			want: head2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseBetweenBefore(tt.args.head, tt.args.n)
			compareNode := tt.want
			fmt.Println("got: ", got.ToArray())
			fmt.Println("compareNode: ", tt.want.ToArray())
			for node := compareNode; node != nil; node = node.Next {
				if got == nil || got.Val != node.Val {
					t.Errorf("reverseBetweenBefore() = %v, want %v", got, tt.want)
					break
				}
				got = got.Next
			}
		})
	}
}

func Test_reverseBetween2(t *testing.T) {
	head1 := link.NewList([]int{1, 2, 3, 4, 5})
	head2 := link.NewList([]int{1, 4, 3, 2, 5})
	head3 := link.NewList([]int{3, 5})
	head4 := link.NewList([]int{5, 3})
	//head2 := link.NewList([]int{4, 3, 2, 1, 5})
	type args struct {
		head *link.ListNode
		m    int
		n    int
	}
	tests := []struct {
		name string
		args args
		want *link.ListNode
	}{
		{
			name: "test1",
			args: args{head: head1, m: 2, n: 4},
			want: head2,
		},
		{
			name: "test2",
			args: args{head: head3, m: 1, n: 2},
			want: head4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseBetween2(tt.args.head, tt.args.m, tt.args.n)
			compareNode := tt.want
			fmt.Println("got: ", got.ToArray())
			fmt.Println("compareNode: ", tt.want.ToArray())
			for node := compareNode; node != nil; node = node.Next {
				if got == nil || got.Val != node.Val {
					t.Errorf("reverseBetween() = %v, want %v", got, tt.want)
					break
				}
				got = got.Next
			}
		})
	}
}

func Test_reverseBetweenBefore2(t *testing.T) {
	head1 := link.NewList([]int{1, 2, 3, 4, 5})
	head2 := link.NewList([]int{4, 3, 2, 1, 5})
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
			args: args{head: head1, n: 4},
			want: head2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseBetweenBefore2(tt.args.head, tt.args.n)
			compareNode := tt.want
			fmt.Println("got: ", got.ToArray())
			fmt.Println("compareNode: ", tt.want.ToArray())
			for node := compareNode; node != nil; node = node.Next {
				if got == nil || got.Val != node.Val {
					t.Errorf("reverseBetweenBefore() = %v, want %v", got, tt.want)
					break
				}
				got = got.Next
			}
		})
	}
}
