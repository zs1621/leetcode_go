package ChapterSearch

import "fmt"

type LNode struct {
	Key  string
	Val  int
	Next *LNode
}

func (n *LNode) hasNext() bool {
	return n.Next != nil
}

type SequentialSearchST struct {
	Head    *LNode
	Size    int
	MaxNode *LNode
}

func (ln *SequentialSearchST) Put(key string) {
	for temp := ln.Head; temp != nil; temp = temp.Next {
		if temp.Key == key {
			temp.Val++
			if temp.Val > ln.MaxNode.Val {
				ln.MaxNode = temp
			}
			return
		}
	}
	ln.Head = &LNode{key, 1, ln.Head}
	ln.Size++
}
func (ln *SequentialSearchST) putval(key string, val int) {
	if ln.Head.Key == "" {
		fmt.Println(ln.Head.Key, "-------")
		ln.Head = &LNode{key, val, nil}
		ln.Size++
	} else {
		for temp := ln.Head; temp != nil; temp = temp.Next {
			if temp.Key == key {
				temp.Val = val
				return
			}
		}
		ln.Head = &LNode{key, val, ln.Head}
		ln.Size++
	}
}

func (ln *SequentialSearchST) Get(key string) int {
	// 找不到返回-1
	if len(key) == 0 {
		return -1
	}
	for temp := ln.Head; temp != nil; temp = temp.Next {
		if temp.Key == key {
			return temp.Val
		}
	}
	return -1
}
