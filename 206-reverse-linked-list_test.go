package leetcode

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/luofeng1/leetcode/link"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		head *link.ListNode
	}
	head1 := link.NewList([]int{1, 2, 3, 4, 5})
	tests := []struct {
		name string
		args args
		want *link.ListNode
	}{
		{
			name: "test1",
			args: args{head1},
			want: head1.Next.Next.Next.Next,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			} else {
				fmt.Println(got.ToArray())
			}
		})
	}
}
