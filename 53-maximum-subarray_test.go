package leetcode

import "testing"

func Test_maxSubArray(t *testing.T) {
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
			args: args{
				nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			},
			want: 6,
		},
		{
			name: "test2",
			args: args{
				nums: []int{-2, -1},
			},
			want: -1,
		},
		{
			name: "test3",
			args: args{
				nums: []int{1, 2, -1, -2, 2, 1, -2, 1},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray(tt.args.nums); got != tt.want {
				t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxSubArray2(t *testing.T) {
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
			args: args{
				nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			},
			want: 6,
		},
		{
			name: "test2",
			args: args{
				nums: []int{-2, -1},
			},
			want: -1,
		},
		{
			name: "test3",
			args: args{
				nums: []int{1, 2, -1, -2, 2, 1, -2, 1},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray2(tt.args.nums); got != tt.want {
				t.Errorf("maxSubArray2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxSubArray3(t *testing.T) {
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
			args: args{
				nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			},
			want: 6,
		},
		{
			name: "test2",
			args: args{
				nums: []int{-2, -1},
			},
			want: -1,
		},
		{
			name: "test3",
			args: args{
				nums: []int{1, 2, -1, -2, 2, 1, -2, 1},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray3(tt.args.nums); got != tt.want {
				t.Errorf("maxSubArray3() = %v, want %v", got, tt.want)
			}
		})
	}
}