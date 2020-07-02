package ChapterSearch


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

	for nt.Current = nt.Root; nt.Current != nil; {
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


func (nt *NodeTree) Select(rank int) string {
	if nt.Size == 0 {
		return ""
	}
	var N, x int
	N = 1
	for nt.Current = nt.Root; nt.Current != nil; {
		if nt.Current.Left == nil {
			x = 0
		} else {
			x = nt.Current.Left.N + 1
		}
		if x + N > rank {
			nt.Current = nt.Current.Left
 		} else if x + N < rank {
			N = x + N
			if nt.Current.Right == nil {
				return nt.Current.Key
			}
			nt.Current = nt.Current.Right
		} else {
			return nt.Current.Key
		}
	}
	return nt.Current.Key
}



//func (nt *NodeTree) Rank(key string) int {
//
//}