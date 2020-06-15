package leetcode

import (
	"reflect"
	"testing"
)

func Test_restoreIpAddresses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test1",
			args: args{"25525511135"},
			want: []string{"255.255.11.135", "255.255.111.35"},
		},
		{
			name: "test2",
			args: args{"010010"},
			want: []string{"0.10.0.10", "0.100.1.0"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := restoreIpAddresses(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("restoreIpAddresses() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_restoreIpAddressesIp2Parts(t *testing.T) {
	type args struct {
		s        string
		position int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test1",
			args: args{"25525511135", 2},
			want: []string{"5", "52"},
		},
		{
			name: "test2",
			args: args{"010", 0},
			want: []string{"0"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := restoreIpAddressesIp2Parts(&tt.args.s, tt.args.position); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("restoreIpAddressesIp2Parts() = %v, want %v", got, tt.want)
			}
		})
	}
}
