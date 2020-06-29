package ChapterSearch

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBST(t *testing.T) {
	x := &NodeTree{nil, 0, nil}
	x.Put("a")
	assert.Equal(t, x.Size, 1)
	assert.Equal(t, x.Root.Value, 1)
	assert.Equal(t, x.Root.N, 0)

	x.Put("a")
	assert.Equal(t, x.Root.Value, 2)
	assert.Equal(t, x.Root.Value, 2)
	assert.Equal(t, x.Root.N, 0)

	x.Put("b")
	assert.Equal(t, x.Root.N, 1)

	x.Put("c")
	assert.Equal(t, x.Root.N, 2)
}
