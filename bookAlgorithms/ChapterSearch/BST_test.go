package ChapterSearch

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"strings"
	"testing"
)
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func initNodeTree() *NodeTree {
	dat, err := ioutil.ReadFile("./tale.txt")
	check(err)
	var x = &NodeTree{nil, 0, nil}
	for _, value := range strings.Split(string(dat), " ") {
		if len(value) < 4 {
			continue
		}
		x.Put(value)
	}
	return x
}
func TestBST(t *testing.T) {
	t.Run("PUT", func(t *testing.T) {
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
	})
	t.Run("GET", func(t *testing.T) {
		x := &NodeTree{nil, 0, nil}
		y := x.Get("c")
		assert.Equal(t, y, -1)
		x.Put("c")
		assert.Equal(t, x.Root.Value, 1)
	})

}


func BenchmarkNodeTree_Get(b *testing.B) {
	b.ResetTimer()
	x := initNodeTree()
	for n := 0; n < b.N; n++ {
		x.Get("xx")
	}
}