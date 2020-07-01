package ChapterSearch

import "fmt"


type Node struct {
	Key   string
	Value int
	Left  *Node
	Right *Node
	N     int
}


type NodeTree struct {
	Root *Node
	Size int
	Current *Node
}


func (nt *NodeTree) Put(data string) {
	temp := &Node{
		Key: data,
		Value: 1,
		Left: nil,
		Right: nil,
		N: 0,
	}
	if nt.Root == nil {
		nt.Root = temp
		nt.Size = 1
		return
	}
	var flagAdd bool
	flagAdd = false
	var nodeList []*Node
	for nt.Current = nt.Root; nt.Current != nil; {
		if nt.Current.Key < data {
			nodeList = append(nodeList, nt.Current)
			if nt.Current.Right == nil {
				nt.Current.Right = temp
				nt.Size++
				flagAdd = true
				break
			}
			nt.Current = nt.Current.Right
		} else if nt.Current.Key > data {
			nodeList = append(nodeList, nt.Current)
			if nt.Current.Left == nil {
				nt.Current.Left = temp
				nt.Size++
				flagAdd = true
				break
			}
			nt.Current = nt.Current.Left
		} else {
			nt.Current.Value++
			break
		}
	}
	if flagAdd {
		for _, nodei := range nodeList {
			nodei.N++
		}
	}
}

func (nt *NodeTree) Get(key string) int {
	for nt.Current != nil {
		if key > nt.Current.Key {
			nt.Current = nt.Current.Right
		} else if key < nt.Current.Key {
			nt.Current = nt.Current.Left
		} else {
			return nt.Current.Value
		}
	}
	return -1
}

func main() {
	x := &Node{
		Key:   "aa",
		Value: 0,
		Left:  nil,
		Right: nil,
		N:     1,
	}
	fmt.Println(x.Key, x.Value)
}
