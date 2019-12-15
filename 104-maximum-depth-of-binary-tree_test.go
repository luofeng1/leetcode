package leetcode

import (
	"github.com/luofeng1/leetcode/tree"
	"testing"
)

func Test_maxDepth(t *testing.T) {
	type args struct {
		root *tree.TreeNode
	}
	tree1 := tree.NewTree([]int{3, 9, 20, tree.INT_MIN, tree.INT_MIN, 15, 7})
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{root: tree1},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.args.root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
