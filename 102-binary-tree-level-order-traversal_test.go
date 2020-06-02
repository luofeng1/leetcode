package leetcode

import (
	"reflect"
	"testing"

	"github.com/luofeng1/leetcode/tree"
)

func Test_levelOrder(t *testing.T) {
	tree1 := tree.NewTree([]int{3, 9, 20, tree.INT_MIN, tree.INT_MIN, 15, 7})
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
			args: args{root: tree1},
			want: [][]int{{3}, {9, 20}, {15, 7}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_levelOrderInOneLevelByBFS(t *testing.T) {
	tree1 := tree.NewTree([]int{3, 9, 20, tree.INT_MIN, tree.INT_MIN, 15, 7})
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
			want: []int{3, 9, 20, 15, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrderInOneLevelByBFS(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrderInOneLevelByBFS() = %v, want %v", got, tt.want)
			}
		})
	}
}
