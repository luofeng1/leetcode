package leetcode

import (
	"reflect"
	"testing"
)

func Test_permute(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test1",
			args: args{[]int{1, 2, 3}},
			want: [][]int{
				{3, 2, 1},
				{2, 3, 1},
				{3, 1, 2},
				{1, 3, 2},
				{2, 1, 3},
				{1, 2, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permute(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_permuteRecursion(t *testing.T) {
	type args struct {
		choiceList []int
		pathList   []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test1",
			args: args{[]int{1, 2, 3}, []int{}},
			want: [][]int{
				{3, 2, 1},
				{2, 3, 1},
				{3, 1, 2},
				{1, 3, 2},
				{2, 1, 3},
				{1, 2, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permuteRecursion(tt.args.choiceList, tt.args.pathList); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permuteRecursion() = %v, want %v", got, tt.want)
			}
		})
	}
}
