package ChapterSearch

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func TestBST(t *testing.T) {
	x := &NodeTree{nil, 0, nil}

	x.Put("c")
	assert.Equal(t, x.Root.Value, 1)
	assert.Equal(t, x.Root.N, 0)

	x.Put("b")
	assert.Equal(t, x.Root.N, 1)
	assert.Nil(t, x.Root.Right)


	x.Put("d")
	assert.Equal(t, x.Root.N, 2)
	assert.Greater(t, x.Root.Right.Key, x.Root.Left.Key)

	x.Put("c")
	assert.Equal(t, x.Root.N, 2)
	assert.Equal(t, x.Root.Value, 2)
}
