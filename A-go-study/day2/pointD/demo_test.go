package main

import "testing"

func Test_add(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name  string
		args  args
		wantA int
	}{
		{name: "test1", args: args{i: 1, j: 2}, wantA: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotA := add(tt.args.i, tt.args.j); gotA != tt.wantA {
				t.Errorf("add() = %v, want %v", gotA, tt.wantA)
			}
		})
	}
}

func Test_c(t *testing.T) {
	c()
}
