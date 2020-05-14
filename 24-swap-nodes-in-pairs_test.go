package leetcode

import (
	"reflect"
	"testing"

	"github.com/luofeng1/leetcode/link"
)

func Test_swapPairs(t *testing.T) {
	head1 := link.NewList([]int{1, 2})
	head2 := link.NewList([]int{2, 1})

	head3 := link.NewList([]int{1, 2, 3, 4})
	head4 := link.NewList([]int{2, 1, 4, 3})
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
			if got := swapPairs(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("swapPairs() = %v, want %v", got, tt.want)
			} else {
				//fmt.Println(got.ToArray())
			}
		})
	}
}
