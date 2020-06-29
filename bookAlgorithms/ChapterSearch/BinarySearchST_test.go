package ChapterSearch

import (
	"gopkg.in/go-playground/assert.v1"
	"testing"
)

func TestBinarySearchST_Get(t *testing.T) {
	type args struct {
		keys []string
		key string
	}
	tests := []struct {
		name   string
		args   args
		want   int
	}{
		{
			"t1",
			args{
				[]string{"a", "b", "c"},
				"bb"},
			2,
		},
		{
			"t2",
			args{
				[]string{"b"},
				"c"},
			1,
		},
		{
			"t3",
			args{
				[]string{"b", "d", "e", "f"},
				"c"},
			1,
		},


	}
	for _, tt := range tests{
		t.Run(tt.name, func(t *testing.T) {
			if got := ArryIndexGet(tt.args.keys, tt.args.key); got != tt.want {
				t.Errorf("get() = %v, want %v", got, tt.want)
			}
		})
	}

}


func TestBinarySearchST_Put(t *testing.T) {
	t.Run("update key", func(t *testing.T) {
		am := &BinarySearchST{
			Keys: []string{},
			Vals: []int{},
			Size: 0,
		}
		am.Put("b")
		assert.Equal(t, am.Size, 1)
		am.Put("b")
		assert.Equal(t, am.Size, 1)

		am.Put("d")
		assert.Equal(t, am.Size, 2)

		am.Put("c")
		assert.Equal(t, am.Size, 3)
		am.Put("c")
		assert.Equal(t, am.Size, 3)
		assert.Equal(t, am.Get("c"), 2)

		am.Put("a")
		assert.Equal(t, am.Size, 4)
		assert.Equal(t, am.Get("a"), 1)

	})
}