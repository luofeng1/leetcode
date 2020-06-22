package leetcode

import (
	"reflect"
	"testing"
)

func Test_exist(t *testing.T) {
	type args struct {
		board [][]byte
		word  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				[][]byte{
					{'A', 'B', 'C', 'E'},
					{'S', 'F', 'C', 'S'},
					{'A', 'D', 'E', 'E'},
				},
				"ABCCED",
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				[][]byte{
					{'A', 'B', 'C', 'E'},
					{'S', 'F', 'C', 'S'},
					{'A', 'D', 'E', 'E'},
				},
				"SEE",
			},
			want: true,
		},
		{
			name: "test3",
			args: args{
				[][]byte{
					{'A', 'B', 'C', 'E'},
					{'S', 'F', 'C', 'S'},
					{'A', 'D', 'E', 'E'},
				},
				"ABCB",
			},
			want: false,
		},
		{
			name: "test4",
			args: args{
				[][]byte{
					{'C', 'A', 'A'},
					{'A', 'A', 'A'},
					{'B', 'C', 'D'},
				},
				"AAB",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exist(tt.args.board, tt.args.word); got != tt.want {
				t.Errorf("exist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_existGetChoices(t *testing.T) {
	type args struct {
		path [][]byte
		x    int
		y    int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test1",
			args: args{
				path: [][]byte{
					{'1', '1'},
					{'1', '1'},
				},
				x: 0,
				y: 0,
			},
			want: [][]int{{1, 0}, {0, 1}},
		},
		{
			name: "test2",
			args: args{
				path: [][]byte{
					{'1', '1', '1'},
					{'1', '1', '1'},
				},
				x: 1,
				y: 1,
			},
			want: [][]int{{0, 1}, {1, 0}, {1, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := existGetChoices(tt.args.path, tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("existGetChoices() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_existDFS(t *testing.T) {
	str1 := "BD"
	str2 := "BC"
	str3 := "AAB"
	str4 := "ABCESEEEFS"

	type args struct {
		choice   [][]int
		path     [][]byte
		word     *string
		position int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				[][]int{{0, 1}},
				[][]byte{{'A', 'B'}, {'C', 'D'}},
				&str1,
				0,
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				[][]int{{0, 1}},
				[][]byte{{'A', 'B'}, {'C', 'D'}},
				&str2,
				0,
			},
			want: false,
		},
		{
			name: "test3",
			args: args{
				[][]int{{1, 1}},
				[][]byte{
					{'C', 'A', 'A'},
					{'A', 'A', 'A'},
					{'B', 'C', 'D'},
				},
				&str3,
				0,
			},
			want: true,
		},
		{
			name: "test4",
			args: args{
				[][]int{{0, 0}},
				[][]byte{
					{'A', 'B', 'C', 'E'},
					{'S', 'F', 'E', 'S'},
					{'A', 'D', 'E', 'E'},
				},
				&str4,
				0,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := existDFS(tt.args.choice, tt.args.path, tt.args.word, tt.args.position); got != tt.want {
				t.Errorf("existDFS() = %v, want %v", got, tt.want)
			}
		})
	}
}
