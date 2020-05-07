package sortx

import (
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func Test_SortArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{5, 3, 2, 1}}, []int{1, 2, 3, 5}},
		{"2", args{[]int{5, 3}}, []int{3, 5}},
		{"3", args{[]int{5}}, []int{5}},
		{"4", args{[]int{5, 1, 1, 2, 0, 0}}, []int{0, 0, 1, 1, 2, 5}},
		{"5", args{[]int{-4, 0, 7, 4, 9, -5, -1, 0, -7, -1}}, []int{-7, -5, -4, -1, -1, 0, 0, 4, 7, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sortArray(tt.args.nums)
			assert.Equal(t, got, tt.want)
		})
	}
}

//
