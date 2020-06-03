package leetcode

import (
	"reflect"
	"testing"

	"github.com/luofeng1/leetcode/tree"
)

func Test_buildTree(t *testing.T) {
	tree1 := tree.NewTree([]int{3, 9, 20, tree.INT_MIN, tree.INT_MIN, 15, 7})
	type args struct {
		preorder []int
		inorder  []int
	}
	tests := []struct {
		name string
		args args
		want *tree.TreeNode
	}{
		{
			name: "test1",
			args: args{[]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}},
			want: tree1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildTree(tt.args.preorder, tt.args.inorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
