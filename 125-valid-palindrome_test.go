package leetcode

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "A man, a plan, a canal: Panama",
			args: args{
				s: "A man, a plan, a canal: Panama",
			},
			want: true,
		},
		{
			name: "1 A man, a plan, a canal: Panama1",
			args: args{
				s: "1 A man, a plan, a canal: Panama1",
			},
			want: true,
		},
		{
			name: "race a car",
			args: args{
				s: "race a car",
			},
			want: false,
		},
		{
			name: "",
			args: args{
				s: "",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidLetterOrDigit(t *testing.T) {
	type args struct {
		c rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "a",
			args: args{c: []rune("a")[0]},
			want: true,
		},
		{
			name: ",",
			args: args{c: []rune(",")[0]},
			want: false,
		},
		{
			name: "A",
			args: args{c: []rune("A")[0]},
			want: true,
		},
		{
			name: "1",
			args: args{c: []rune("1")[0]},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidLetterOrDigit(tt.args.c); got != tt.want {
				t.Errorf("isValidLetterOrDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}
