package leetcode

import (
	"reflect"
	"testing"

	"github.com/luofeng1/leetcode/tree"
)

func Test_lowestCommonAncestor(t *testing.T) {
	tree1 := tree.NewTree([]int{3, 5, 1, 6, 2, 0, 8, tree.INT_MIN, tree.INT_MIN, 7, 4})
	type args struct {
		root *tree.TreeNode
		p    *tree.TreeNode
		q    *tree.TreeNode
	}
	tests := []struct {
		name string
		args args
		want *tree.TreeNode
	}{
		{
			name: "test1",
			args: args{tree1, tree1.Left, tree1.Right},
			want: tree1,
		},
		{
			name: "test2",
			args: args{tree1, tree1.Left, tree1},
			want: tree1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestor(tt.args.root, tt.args.p, tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got, tt.want)
			}
		})
	}
}
