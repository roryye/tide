package main

import (
	"reflect"
	"testing"
)

func Test_findIndex(t *testing.T) {
	type args struct {
		target int
		nums   []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case1",
			args: args{
				nums:   []int{1, 2, 2, 3, 4, 7, 9, 9},
				target: 2,
			},
			want: []int{1, 2},
		},
		{
			name: "case2",
			args: args{
				nums:   []int{1, 2, 2, 3, 4, 7, 9, 9},
				target: 5,
			},
			want: nil,
		},
		{
			name: "case3",
			args: args{
				nums:   []int{1, 2, 2, 3, 4, 7, 9, 9},
				target: 9,
			},
			want: []int{6, 7},
		},
		{
			name: "case4",
			args: args{
				nums:   []int{1, 2, 2, 3, 4, 7, 9, 9},
				target: 11,
			},
			want: nil,
		},
		{
			name: "case5",
			args: args{
				nums:   []int{1},
				target: 1,
			},
			want: []int{0},
		},
		{
			name: "case6",
			args: args{
				nums:   []int{},
				target: 1,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findIndex(tt.args.target, tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
