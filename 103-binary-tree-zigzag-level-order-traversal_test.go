package leetcode

import (
	"reflect"
	"testing"

	"github.com/luofeng1/leetcode/tree"
)

func Test_zigzagLevelOrder(t *testing.T) {
	tree1 := tree.NewTree([]int{3, 9, 20, tree.INT_MIN, tree.INT_MIN, 15, 7})
	tree2 := tree.NewTree([]int{1, 2, 3, 4, tree.INT_MIN, tree.INT_MIN, 5})
	type args struct {
		root *tree.TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test1",
			args: args{tree1},
			want: [][]int{
				{3},
				{20, 9},
				{15, 7},
			},
		},
		{
			name: "test2",
			args: args{tree2},
			want: [][]int{
				{1},
				{3, 2},
				{4, 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := zigzagLevelOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("zigzagLevelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
