package leetcode

import (
	"testing"

	"github.com/luofeng1/leetcode/tree"
)

func Test_isValidBST(t *testing.T) {
	tree1 := tree.NewTree([]int{5, 1, 4, tree.INT_MIN, tree.INT_MIN, 3, 6})
	tree2 := tree.NewTree([]int{2, 1, 3})
	tree3 := tree.NewTree([]int{5, 1, 7, tree.INT_MIN, tree.INT_MIN, 6, 8})
	type args struct {
		root *tree.TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{tree1},
			want: false,
		},
		{
			name: "test2",
			args: args{tree2},
			want: true,
		},
		{
			name: "test3",
			args: args{tree3},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.args.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidBST_2(t *testing.T) {
	tree1 := tree.NewTree([]int{5, 1, 4, tree.INT_MIN, tree.INT_MIN, 3, 6})
	tree2 := tree.NewTree([]int{2, 1, 3})
	tree3 := tree.NewTree([]int{5, 1, 7, tree.INT_MIN, tree.INT_MIN, 6, 8})
	type args struct {
		root *tree.TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{tree1},
			want: false,
		},
		{
			name: "test2",
			args: args{tree2},
			want: true,
		},
		{
			name: "test3",
			args: args{tree3},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST_2(tt.args.root); got != tt.want {
				t.Errorf("isValidBST_2() = %v, want %v", got, tt.want)
			}
		})
	}
}