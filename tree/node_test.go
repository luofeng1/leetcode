package tree

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func PrintTest(i interface{}) {
	bytes, _ := json.Marshal(i)
	fmt.Println(string(bytes))
}

func TestNewTree(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "[3,9,20]",
			args: args{
				arr: []int{3, 9, 20},
			},
			want: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   9,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   20,
					Left:  nil,
					Right: nil,
				},
			},
		},
		{
			name: "[3,INT_MIN,20]",
			args: args{
				arr: []int{3, INT_MIN, 20},
			},
			want: &TreeNode{
				Val:  3,
				Left: nil,
				Right: &TreeNode{
					Val:   20,
					Left:  nil,
					Right: nil,
				},
			},
		},
		{
			name: "[3,9,20,8,9,10]",
			args: args{
				arr: []int{3, 9, 20, 8, 9, 10},
			},
			want: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
					Left: &TreeNode{
						Val:   8,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   9,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val:   10,
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
			},
		},
		{
			name: "[3,9,20,null,null,5,7]",
			args: args{
				arr: []int{3, 9, 20, ^int(^uint(0) >> 1), ^int(^uint(0) >> 1), 5, 7},
			},
			want: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val:   5,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   7,
						Left:  nil,
						Right: nil,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTree(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				PrintTest(got)
				PrintTest(got.Left)
				t.Errorf("NewTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getParentNode(t *testing.T) {
	type args struct {
		parentsNode []*TreeNode
		depth       int
		index       int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "test1",
			args: args{
				parentsNode: []*TreeNode{
					&TreeNode{Val: 100},
					&TreeNode{Val: 200},
					&TreeNode{Val: 300},
					&TreeNode{Val: 400},
				},
				depth: 4,
				index: 10,
			},
			want: &TreeNode{Val: 200},
		},
		{
			name: "test2",
			args: args{
				parentsNode: []*TreeNode{
					&TreeNode{Val: 100},
					&TreeNode{Val: 200},
				},
				depth: 3,
				index: 5,
			},
			want: &TreeNode{Val: 200},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getParentNode(tt.args.parentsNode, tt.args.depth, tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getParentNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isNewDepth(t *testing.T) {
	type args struct {
		index int
		depth int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				index: 1,
				depth: 1,
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				index: 3,
				depth: 2,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNewDepth(tt.args.index, tt.args.depth); got != tt.want {
				t.Errorf("newDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
