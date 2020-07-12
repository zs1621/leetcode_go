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

	t.Run("Min", func(t *testing.T) {
		x := &NodeTree{nil, 0, nil}
		x.Put("x")
		assert.Equal(t, Min(x.Root), x.Root)

		x.Put("m")
		assert.Equal(t, Min(x.Root), x.Root.Left)

	})

	t.Run("SELECT", func(t *testing.T) {
		x := &NodeTree{nil, 0, nil}
		y := x.Select(1)
		assert.Equal(t, "", y)

		x.Put("d")
		assert.Equal(t, "d", x.Select(1))

		x.Put("c")
		assert.Equal(t, "c", x.Select(1))

		x.Put("e")
		assert.Equal(t, "e", x.Select(3))

		x.Put("b")
		assert.Equal(t, "b", x.Select(1))

		x.Put("a")
		assert.Equal(t, "a", x.Select(1))

	})

	t.Run("Rank", func(t *testing.T) {
		x := &NodeTree{nil, 0, nil}
		x.Put("c")
		assert.Equal(t, 1, x.Rank("c"))

		x.Put("d")
		assert.Equal(t, 2, x.Rank("d"))

		x.Put("a")
		assert.Equal(t, 1, x.Rank("a"))
		assert.Equal(t, 3, x.Rank("d"))

		x.Put("b")
		assert.Equal(t, 2, x.Rank("b"))

		x.Put("x")
		assert.Equal(t, 5, x.Rank("x"))
		assert.Equal(t, -1, x.Rank("e"))

		x.Put("e")
		assert.Equal(t, 5, x.Rank("e"))
		assert.Equal(t, -1, x.Rank("t"))
		assert.Equal(t, 2, x.Rank("b"))
	})
	//
	//t.Run("Delete", func(t *testing.T) {
	//
	//})
	t.Run("Floor-Ceil", func(t *testing.T) {
		x := &NodeTree{nil, 0, nil}
		x.Put("y")
		x.Put("k")
		x.Put("q")
		x.Put("s")
		y := x.Floor("j")
		assert.Equal(t, "", y)

		y = x.Floor("u")
		assert.Equal(t, "s", y)

		y = x.Ceil("l")
		assert.Equal(t, "q", y)

		y = x.Ceil("z")
		assert.Equal(t, "", y)

		y = x.Ceil("r")
		assert.Equal(t, "s", y)
	})
}

func BenchmarkNodeTree_Get(b *testing.B) {
	b.ResetTimer()
	x := initNodeTree()
	for n := 0; n < b.N; n++ {
		x.Get("xx")
	}
}
