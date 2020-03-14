package leetcode

import "testing"

func Test_findTargetSumWays(t *testing.T) {
	type args struct {
		nums []int
		S    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				nums: []int{1, 1, 1, 1, 1},
				S:    3,
			},
			want: 5,
		},
		{
			name: "test2",
			args: args{
				nums: []int{1},
				S:    1,
			},
			want: 1,
		},
		{
			name: "test3",
			args: args{
				nums: []int{1, 0},
				S:    1,
			},
			want: 2,
		},
		{
			name: "test4",
			args: args{
				nums: []int{4, 1, 7, 1, 8, 7, 8, 7, 7, 4},
				S:    4,
			},
			want: 22,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTargetSumWays(tt.args.nums, tt.args.S); got != tt.want {
				t.Errorf("findTargetSumWays() = %v, want %v", got, tt.want)
			}
		})
	}
}
