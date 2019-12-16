package leetcode

import (
	"reflect"
	"testing"

	"github.com/luofeng1/leetcode/link"
)

func Test_addTwoNumbers(t *testing.T) {
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
			name: "342 + 465 = 807",
			args: args{
				l1: link.NewList([]int{2, 4, 3}),
				l2: link.NewList([]int{5, 6, 4}),
			},
			want: link.NewList([]int{7, 0, 8}),
		},
		{
			name: "642 + 465 = 1107",
			args: args{
				l1: link.NewList([]int{2, 4, 6}),
				l2: link.NewList([]int{5, 6, 4}),
			},
			want: link.NewList([]int{7, 0, 1, 1}),
		},
		{
			name: "10 + 0 = 10",
			args: args{
				l1: link.NewList([]int{0, 1}),
				l2: link.NewList([]int{}),
			},
			want: link.NewList([]int{0, 1}),
		},
		{
			name: "10 + 210 = 220",
			args: args{
				l1: link.NewList([]int{0, 1}),
				l2: link.NewList([]int{0, 1, 2}),
			},
			want: link.NewList([]int{0, 2, 2}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				PrintTest(got)
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
