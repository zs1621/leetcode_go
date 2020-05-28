package link

type Node struct {
	Key  string
	Val  int
	Next *Node
}

func NewNode() *Node {
	return &Node{"", 0, nil}
}
func (n *Node) hasNext() bool {
	return n.Next != nil
}

type LinkNode struct {
	Head *Node
	Size int
}


func (ln *LinkNode) put(key string, val int) {
	if ln.Head.Key == "" {
		ln.Head = &Node{key, val, nil}
		ln.Size ++
	} else {
		temp := ln.Head
		for temp.hasNext() {
			temp = temp.Next
		}
		temp.Next = &Node{key, val, nil}
		ln.Size ++
	}
}

func (ln *LinkNode) get(key string) int {
	return 1
}



