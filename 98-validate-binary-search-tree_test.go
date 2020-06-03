package leetcode

import (
	"fmt"
	"testing"

	"github.com/luofeng1/leetcode/tree"
)

func Test_isValidBST(t *testing.T) {
	tree1 := tree.NewTree([]int{5, 1, 4, tree.INT_MIN, tree.INT_MIN, 3, 6})
	tree2 := tree.NewTree([]int{2, 1, 3})
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.args.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidBSTByMiddle(t *testing.T) {
	tree1 := tree.NewTree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	arr1 := &[]int{}
	type args struct {
		root *tree.TreeNode
		arr  *[]int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{tree1, arr1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isValidBSTByMiddle(tt.args.root, tt.args.arr)
			fmt.Println("中序结果:", arr1)
		})
	}
}
