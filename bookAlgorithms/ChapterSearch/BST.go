package ChapterSearch

type Node struct {
	Key   string
	Value int
	Left  *Node
	Right *Node
	N     int
}

type NodeTree struct {
	Root    *Node
	Size    int
	Current *Node
}

func (nt *NodeTree) Put(data string) {
	temp := &Node{
		Key:   data,
		Value: 1,
		Left:  nil,
		Right: nil,
		N:     0,
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
		if x+N > rank {
			nt.Current = nt.Current.Left
		} else if x+N < rank {
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

func (nt *NodeTree) Rank(key string) int {
	if nt.Size == 0 {
		return -1
	}
	var N int
	N = 0
	for nt.Current = nt.Root; nt.Current != nil; {
		if nt.Current.Key > key {
			nt.Current = nt.Current.Left
		} else if nt.Current.Key < key {
			if nt.Current.Right != nil {
				N = N + nt.Current.N - nt.Current.Right.N
			} else {
				N = N + nt.Current.N
			}
			nt.Current = nt.Current.Right
		} else {
			break
		}
	}
	if nt.Current == nil {
		return -1
	}
	if nt.Current.Left != nil {
		return N + nt.Current.Left.N
	} else {
		return N + 1
	}
}

func Min(n *Node) *Node {
	if n == nil {
		return nil
	}
	if n.Left == nil {
		return n
	}
	return Min(n.Left)
}

func Max(n *Node) *Node {
	if n.Right == nil {
		return n
	}
	return Max(n.Right)
}

func (nt *NodeTree) Min() string {
	return Min(nt.Root).Key
}
func (nt *NodeTree) Max() string {
	return Max(nt.Root).Key
}


func (nt *NodeTree) Floor(key string) string {
	nt.Current = nt.Root
	lastKey := nt.Current.Key
	for nt.Current != nil {
		if nt.Current.Key > key {
			lastKey = nt.Current.Key
			nt.Current = nt.Current.Left
		} else if nt.Current.Key < key {
			if nt.Current.Right != nil {
				if key < Min(nt.Current.Right).Key {
					return nt.Current.Key
				}
			}
			lastKey = nt.Current.Key
			nt.Current = nt.Current.Right
		} else {
			return nt.Current.Key
		}
	}
	if lastKey > key {
		return ""
	}
	return lastKey
}


func (nt *NodeTree) Ceil(key string) string {
	for nt.Max() < key {
		return ""
	}
	nt.Current = nt.Root
	lastKey := nt.Current.Key
	for nt.Current != nil {
		if nt.Current.Key > key {
			if nt.Current.Left != nil {
				if Max(nt.Current.Left).Key < key {
					return nt.Current.Key
				}
			}
			lastKey = nt.Current.Key
			nt.Current = nt.Current.Left
		} else if nt.Current.Key < key {
			if nt.Current.Right != nil {
				if Max(nt.Current.Right).Key < key {
					return ""
				}
			}
			lastKey = nt.Current.Key
			nt.Current = nt.Current.Right
		} else {
			return key
		}
	}
	if lastKey < key {
		return ""
	}
	return lastKey
}

func Range(n *Node,a []string, lo string, hi string) []string{
	if n == nil {
		return a
	}
	if n.Key > lo {
		a = Range(n.Left, a, lo, hi)
	}
	if  lo <= n.Key  && n.Key <= hi {
		a = append(a, n.Key)
	}
	if n.Key < hi {
		a = Range(n.Right, a, lo, hi)
	}
	return a
}

func (nt *NodeTree) Range(lo string, hi string) []string {
	a := []string{}
	return Range(nt.Root, a, lo, hi)
}

func (nt *NodeTree) Delete(key string) {
	for nt.Current = nt.Root; nt.Current != nil; {
		if nt.Current.Key > key {

		}
	}
}
