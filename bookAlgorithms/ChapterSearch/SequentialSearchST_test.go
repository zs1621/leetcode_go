package ChapterSearch

import (
	"gopkg.in/go-playground/assert.v1"
	"testing"
)

func TestSequentialSearchST_Put(t *testing.T) {
	type fields struct {
		Head *LNode
		Size int
	}
	type args struct {
		key string
		val int
	}
	type result struct{
		size int
		hasNext bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want result
	}{
		{
			"not first insert",
			fields{&LNode{"yy", 1, nil}, 1},
			args{"xx", 1},
			result{2, true},
		},
		{
			"first insert",
			fields{&LNode{"", 0, nil}, 0},
			args{"xx", 1},
			result{1, false},
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ln := &SequentialSearchST{
				Head: tt.fields.Head,
				Size: tt.fields.Size,
			}
			ln.putval(tt.args.key, tt.args.val)
			assert.Equal(t, ln.Size, tt.want.size)
			assert.Equal(t, ln.Head.hasNext(), tt.want.hasNext)
		})
	}
}