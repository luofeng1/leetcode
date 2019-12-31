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
				nums: []int{25, 33, 27, 23, 46, 16, 10, 27, 33, 2, 12, 2, 29, 44, 49, 40, 32, 46, 7, 50},
				S:    4,
			},
			want: 1,
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
