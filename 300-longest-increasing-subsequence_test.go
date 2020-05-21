package leetcode

import "testing"

func Test_lengthOfLIS(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{nums: []int{10, 9, 2, 5}},
			want: 2,
		},
		{
			name: "test2",
			args: args{nums: []int{1, 3, 6, 7, 9, 4, 10, 5, 6}},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLIS(tt.args.nums); got != tt.want {
				t.Errorf("lengthOfLIS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lengthOfLIS2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{nums: []int{10, 9, 2, 5}},
			want: 2,
		},
		{
			name: "test2",
			args: args{nums: []int{1, 3, 6, 7, 9, 4, 10, 5, 6}},
			want: 6,
		},
		{
			name: "test3",
			args: args{nums: []int{}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLIS2(tt.args.nums); got != tt.want {
				t.Errorf("lengthOfLIS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lengthOfLISQuery(t *testing.T) {
	type args struct {
		q   []int
		val int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				q:   []int{2, 3, 4, 56, 78},
				val: 4,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLISQuery(tt.args.q, tt.args.val); got != tt.want {
				t.Errorf("lengthOfLISQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}
