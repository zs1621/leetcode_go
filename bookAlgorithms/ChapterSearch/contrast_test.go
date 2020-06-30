package ChapterSearch

import (
	"io/ioutil"
	"strings"
	"testing"
	"time"
)

func TestSTContrast_Put(t *testing.T) {

	t.Run("无序链表-Put", func(t *testing.T) {
		var x *SequentialSearchST
		x = &SequentialSearchST{nil, 0, nil}
		x.Head = &LNode{"", 0, nil}
		x.MaxNode = x.Head

		t1 := time.Now()
		dat, err := ioutil.ReadFile("./tale.txt")
		check(err)
		for _, value := range strings.Split(string(dat), " ") {
			if len(value) < 4 {
				continue
			}
			x.Put(value)
		}
		elapsed := time.Since(t1)
		t.Logf("SequentialSearchST put use time: %s, size: %d", elapsed, x.Size)
	})

	t.Run("有序数组key:value - Put", func(t *testing.T) {
		var x *BinarySearchST
		x = &BinarySearchST{
			Keys: []string{},
			Vals: []int{},
			Size: 0,
		}
		t1 := time.Now()
		dat, err := ioutil.ReadFile("./tale.txt")
		check(err)
		for _, value := range strings.Split(string(dat), " ") {
			if len(value) < 4 {
				continue
			}
			x.Put(value)
		}
		elapsed := time.Since(t1)
		t.Logf("BinarySearchST put use time: %s, size: %d", elapsed, x.Size)
	})

	t.Run("二叉搜索树-Put", func(t *testing.T) {
		t1 := time.Now()
		dat, err := ioutil.ReadFile("./tale.txt")
		check(err)
		var x = &NodeTree{nil, 0, nil}
		for _, value := range strings.Split(string(dat), " ") {
			if len(value) < 4 {
				continue
			}
			x.Put(value)
		}
		elapsed := time.Since(t1)
		t.Logf("put use time: %s, size: %d", elapsed, x.Size)
	})
}
