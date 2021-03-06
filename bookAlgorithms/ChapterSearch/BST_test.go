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

	t.Run("Range", func(t *testing.T) {
		x := &NodeTree{nil, 0, nil}
		x.Put("y")
		x.Put("k")
		x.Put("q")
		x.Put("s")
		x.Put("m")
		x.Put("x")
		y := x.Range("r", "z")
		assert.Equal(t, []string{"s", "x",  "y"}, y)
		x = &NodeTree{nil, 0, nil}
		y = x.Range("a", "x")
		assert.Equal(t, []string{}, y)
	})

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

	t.Run("Delete", func(t *testing.T) {
		x := &NodeTree{nil, 0, nil}
		x.Put("y")
		x.Delete("y")
		assert.Equal(t, 0, x.Size)

		x.Put("y")
		x.Put("k")
		x.Delete("y")
		assert.Equal(t, 1, x.Size)
		assert.Equal(t, "k", x.Root.Key)

		x = &NodeTree{nil, 0, nil}
		x.Put("y")
		x.Put("z")
		x.Delete("y")
		assert.Equal(t, 1, x.Size)
		assert.Equal(t, "z", x.Root.Key)


		x = &NodeTree{nil, 0, nil}
		x.Put("y")
		x.Put("z")
		x.Put("x")
		x.Delete("y")
		assert.Equal(t, 2, x.Size)
		assert.Equal(t, "z", x.Root.Key)
		assert.Equal(t, "x", x.Root.Left.Key)
		assert.Equal(t, x.Root.Left.Right, x.Root.Right)


		x = &NodeTree{nil, 0, nil}
		x.Put("o")
		x.Put("s")
		x.Put("p")
		x.Put("w")
		x.Put("t")
		x.Put("c")
		assert.Equal(t, 1, x.Root.Right.Right.N)
		x.Delete("s")
		assert.Equal(t, 5, x.Size)
		assert.Equal(t, "t", x.Root.Right.Key)
		assert.Equal(t, 0, x.Root.Right.Right.N)
		assert.Equal(t, 4, x.Root.N)
	})
}

func BenchmarkNodeTree_Get(b *testing.B) {
	b.ResetTimer()
	x := initNodeTree()
	for n := 0; n < b.N; n++ {
		x.Get("xx")
	}
}
