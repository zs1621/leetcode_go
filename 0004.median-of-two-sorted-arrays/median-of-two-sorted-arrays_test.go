package problem1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []int
	two []int
}

type ans struct {
	one float64
}

type question struct {
	p para
	a ans
}

func Test_OK(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: []int{1, 3, 4, 4},
				two: []int{0, 2, 3, 3},
			},
			a: ans{
				one: 3,
			},
		},
		question{
			p: para{
				one: []int{1, 2, 3},
				two: []int{3, 5, 6},
			},
			a: ans{
				one: 3,
			},
		},
		question{
			p: para{
				one: []int{1, 2},
				two: []int{},
			},
			a: ans{
				one: 1.5,
			},
		},
		question{
			p: para{
				one: []int{},
				two: []int{1, 2},
			},
			a: ans{
				one: 1.5,
			},
		},
		question{
			p: para{
				one: []int{1, 2, 4, 6},
				two: []int{2, 3, 4},
			},
			a: ans{
				one: 3,
			},
		},
		question{
			p: para{
				one: []int{1},
				two: []int{1},
			},
			a: ans{
				one: 1,
			},
		},
		question{
			p: para{
				one: []int{2},
				two: []int{1, 3, 4},
			},
			a: ans{
				one: 2.5,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, FindMedianSortedArrays(p.one, p.two), "输入:%v", p)
	}
}
