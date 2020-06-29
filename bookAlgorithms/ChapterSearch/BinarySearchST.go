package ChapterSearch



type BinarySearchST struct {
	Keys []string
	Vals []int
	Size int
}

func ArryIndexGet(keys []string, key string) int {
	start := 0
	end := len(keys) - 1
	for start <= end {
		mid := start + (end - start) / 2
		if keys[mid] > key {
			end = mid - 1
		} else if keys[mid] < key {
			start = mid + 1
		} else {
			return mid
		}
	}
	return start
}
func (ma *BinarySearchST) Put(key string) {
	index := ArryIndexGet(ma.Keys, key)
	lenOrigin := len(ma.Keys)
	if index > lenOrigin - 1  || (index <= lenOrigin - 1 && ma.Keys[index] != key)  {
		ma.Keys = append(ma.Keys, "")
		ma.Vals = append(ma.Vals, 0)
	}
	if ma.Keys[index] == key {
		ma.Vals[index] = ma.Vals[index] + 1
		return
	} else {
		for x := ma.Size; x > index; x-- {
			ma.Keys[x] = ma.Keys[x-1]
			ma.Vals[x] = ma.Vals[x-1]
		}
	}
	ma.Vals[index] = 1
	ma.Keys[index] = key
	ma.Size ++
}

func (ma *BinarySearchST) Get(key string) int {
	index := ArryIndexGet(ma.Keys, key)
	if ma.Keys[index] != key {
		return -1
	} else {
		return ma.Vals[index]
	}
}