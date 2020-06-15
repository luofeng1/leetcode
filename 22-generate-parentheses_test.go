package leetcode

import (
	"reflect"
	"testing"
)

func Test_generateParenthesis(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test1",
			args: args{3},
			want: []string{
				"((()))",
				"(()())",
				"(())()",
				"()(())",
				"()()()",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateParenthesis(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateParenthesis() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateParenthesisGetChoiceList(t *testing.T) {
	type args struct {
		remainderLevel int
		pathList       []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "test1",
			args: args{6, []byte{}},
			want: []byte{'('},
		},
		{
			name: "test2",
			args: args{5, []byte{'('}},
			want: []byte{'(', ')'},
		},
		{
			name: "test3",
			args: args{3, []byte{'(', '(', ')'}},
			want: []byte{'(', ')'},
		},
		{
			name: "test4",
			args: args{3, []byte{'(', '(', '('}},
			want: []byte{')'},
		},
		{
			name: "test45",
			args: args{2, []byte{'(', '(', ')', '('}},
			want: []byte{')'},
		},
		{
			name: "test6",
			args: args{2, []byte{'(', '(', ')', ')'}},
			want: []byte{'('},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateParenthesisGetChoiceList(tt.args.remainderLevel, tt.args.pathList); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateParenthesisGetChoiceList() = %v, want %v", got, tt.want)
			}
		})
	}
}
