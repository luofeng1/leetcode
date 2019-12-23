package leetcode

import "testing"

func Test_uniquePaths(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				m: 7,
				n: 3,
			},
			want: 28,
		},
		{
			name: "test2",
			args: args{
				m: 3,
				n: 7,
			},
			want: 28,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePaths(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("uniquePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_uniquePaths2(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				m: 7,
				n: 3,
			},
			want: 28,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePaths2(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("uniquePaths2() = %v, want %v", got, tt.want)
			}
		})
	}
}
