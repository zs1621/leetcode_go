package basic

import "testing"

func Test_bSearch(t *testing.T) {
	type args struct {
		nums  []int
		value int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 2, 3, 9}, 10}, 4},
		{"2", args{[]int{1, 2, 3, 9}, 1}, 0},
		{"3", args{[]int{1, 2}, 2}, 1},
		{"4", args{[]int{1, 2}, 9}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchInsert(tt.args.nums, tt.args.value); got != tt.want {
				t.Errorf("bSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
