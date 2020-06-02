package leetcode

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/luofeng1/leetcode/tree"
)

func Test_rightSideView(t *testing.T) {
	tree1 := tree.NewTree([]int{1, 2, 3, tree.INT_MIN, 5, tree.INT_MIN, 4})
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
			args: args{root: tree1},
			want: []int{1, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rightSideView(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rightSideView() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rightSideViewByBFS(t *testing.T) {
	tree1 := tree.NewTree([]int{1, 2, 3, tree.INT_MIN, 5, tree.INT_MIN, 4})
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
			args: args{root: tree1},
			want: []int{1, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rightSideViewByBFS(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rightSideViewByBFS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rightSideViewByDFS(t *testing.T) {
	tree1 := tree.NewTree([]int{1, 2, 3, tree.INT_MIN, 5, tree.INT_MIN, 4})
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
			args: args{root: tree1},
			want: []int{1, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rightSideViewByDFS(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rightSideViewByDFS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rightSideViewByDFSPrint(t *testing.T) {
	tree1 := tree.NewTree([]int{1, 2, 3, tree.INT_MIN, 5, tree.INT_MIN, 4})
	type args struct {
		root  *tree.TreeNode
		depth int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{tree1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rightSideViewByDFSPrint(tt.args.root, tt.args.depth)
		})
	}
}

func Test_rightSideViewByDFSInLast(t *testing.T) {
	tree1 := tree.NewTree([]int{1, 2, 3, tree.INT_MIN, 5, tree.INT_MIN, 4})
	result1 := []int{}
	type args struct {
		root         *tree.TreeNode
		depth        int
		depthLastVal *[]int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{tree1, 0, &result1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rightSideViewByDFSInLast(tt.args.root, tt.args.depth, tt.args.depthLastVal)
			fmt.Println("got: ", tt.args.depthLastVal)
		})
	}
}
