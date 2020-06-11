package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_decodeString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{"3[a]2[bc]"},
			want: "aaabcbc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decodeString(tt.args.s); got != tt.want {
				t.Errorf("decodeString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_decodeStringDFS1(t *testing.T) {
	position1 := 0
	position2 := 0
	position3 := 0
	position4 := 0
	position5 := 0
	type args struct {
		s        []byte
		position *int
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "test1",
			args: args{[]byte("swqe"), &position1},
			want: []byte("swqe"),
		},
		{
			name: "test2",
			args: args{[]byte("3[a]2[bc]"), &position2},
			want: []byte("aaabcbc"),
		},
		{
			name: "test3",
			args: args{[]byte("3[a2[c]]"), &position3},
			want: []byte("accaccacc"),
		},
		{
			name: "test4",
			args: args{[]byte("2[abc]3[cd]ef"), &position4},
			want: []byte("abcabccdcdcdef"),
		},
		{
			name: "test5",
			args: args{[]byte("10[a]"), &position5},
			want: []byte("aaaaaaaaaa"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decodeStringDFS(tt.args.s, tt.args.position); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decodeStringDFS() = %v, want %v", string(got), string(tt.want))
			} else {
				fmt.Printf(string(got))
			}
		})
	}
}
