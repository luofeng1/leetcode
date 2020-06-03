package leetcode

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/luofeng1/leetcode/tree"
)

func Test_inorderTraversal(t *testing.T) {
	tree1 := tree.NewTree([]int{1, tree.INT_MIN, 2, tree.INT_MIN, tree.INT_MIN, 3})
	tree2 := tree.NewTree([]int{1, 2, 3, 4, 5, 6, 7, tree.INT_MIN, 8, 9})
	type args struct {
		root *tree.TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{tree1},
			want: []int{1, 3, 2},
		},
		{
			name: "test2",
			args: args{tree2},
			want: []int{4, 8, 2, 9, 5, 1, 6, 3, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderTraversal() = %v, want %v", got, tt.want)
			} else {
				fmt.Println("got:", got)
			}
		})
	}
}

func Test_inorderTraversalByRecursion(t *testing.T) {
	tree1 := tree.NewTree([]int{1, tree.INT_MIN, 2, tree.INT_MIN, tree.INT_MIN, 3})
	type args struct {
		root *tree.TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{tree1},
			want: []int{1, 3, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inorderTraversalByRecursion(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderTraversalByRecursion() = %v, want %v", got, tt.want)
			} else {
				fmt.Println("got:", got)
			}
		})
	}
}

func Test_inorderTraversalByColorInStack(t *testing.T) {
	tree2 := tree.NewTree([]int{1, 2, 3, 4, 5, 6, 7, tree.INT_MIN, 8, 9})
	type args struct {
		root *tree.TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test2",
			args: args{tree2},
			want: []int{4, 8, 2, 9, 5, 1, 6, 3, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inorderTraversalByColorInStack(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderTraversalByColorInStack() = %v, want %v", got, tt.want)
			}
		})
	}
}
