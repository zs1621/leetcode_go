package sortx

func sortArrayBubble(nums []int) []int {
	var (
		i        int
		j        int
		valueMin int
		indexMin int
	)
	for i = 0; i < len(nums); i++ {
		valueMin = nums[i]
		indexMin = i
		for j = i; j < len(nums); j++ {
			if valueMin > nums[j] {
				valueMin = nums[j]
				indexMin = j
			}
		}
		if indexMin != j {
			nums[indexMin] = nums[i]
			nums[i] = valueMin
		}
	}
	return nums
}

func sortArrayInsertImprove(nums []int) []int {
	var (
		i int
		j int
		n int
	)
	n = len(nums)
	for i = 0; i < n; i++ {
		temp := nums[i]
		for j = i; j > 0 && nums[j-1] > temp; j-- {
			nums[j] = nums[j-1]
		}
		nums[j] = temp
	}
	return nums
}

// 希尔排序
func sortArrayShell(nums []int) []int {
	var (
		N int
		i int
		j int
		h int
	)
	N = len(nums)
	if N == 1 {
		return nums
	}
	h = 1
	for h < N/3 {
		h = 3*h + 1
	}
	for h >= 1 {
		for i = h; i < N; i++ {
			temp := nums[i]
			for j = i; j > h-1 && temp < nums[j-h]; j -= h {
				nums[j] = nums[j-h]
			}
			nums[j] = temp
		}

		h = h / 3
	}
	return nums
}

// 归并排序 --- start-------
const InsertSortThreshold = 7

var tempArray []int

func sortInsert(nums []int, start int, end int) {
	var (
		i int
		j int
	)
	for i = start; i < end+1; i++ {
		temp := nums[i]
		for j = i; j > start && nums[j-1] > temp; j-- {
			nums[j] = nums[j-1]
		}
		nums[j] = temp
	}

}

func sortx(nums []int, low int, high int) {
	var middle int
	if high-low <= InsertSortThreshold {
		sortInsert(nums, low, high)
		return
	}
	if low < high {
		middle = low + (high-low)/2
		sortx(nums, low, middle)
		sortx(nums, middle+1, high)
		mergeSort(nums, low, middle, high)
	}
}

func mergeSort(nums []int, low int, middle int, high int) {
	// 归并两有序数组 [low:middle] [middle+1:high]
	if nums[middle] < nums[middle+1] {
		return
	}
	for key, value := range nums[low : high+1] {
		tempArray[low+key] = value
	}
	var (
		i        int
		j        int
		valuelen int
	)
	i = low
	j = middle + 1
	valuelen = low
	for ; valuelen < high+1; valuelen++ {
		if i > middle {
			nums[valuelen] = tempArray[j]
			j++
		} else if j > high {
			nums[valuelen] = tempArray[i]
			i++
		} else if tempArray[i] <= tempArray[j] {
			nums[valuelen] = tempArray[i]
			i++
		} else {
			nums[valuelen] = tempArray[j]
			j++
		}
	}
}

func sortArrayMerge(nums []int) []int {
	// 归并排序入口 -> 28ms 6.4M
	tempArray = make([]int, len(nums))
	sortx(nums, 0, len(nums)-1)
	return nums
}
// 归并排序 --- end -------



// 快速排序 ----- start -------
func partition(nums []int, lo int, hi int) int {
	var pivo int
	var (
		i int
		j int
		pivoIndex int
	)
	i = lo
	j = hi
	pivoIndex = lo
	pivo = nums[pivoIndex]
	for true {
		for nums[i] < pivo {
			i++
			if i == hi {
				break
			}
		}
		for nums[j] > pivo {
			j--
			if j == lo {
				break
			}
		}
		if i >= j {
			break
		}
		// swap
		temp := nums[j]
		nums[j] = nums[i]
		nums[i] = temp
		i ++
		j --
	}
	//tempx := nums[j]
	//nums[j] = nums[lo]
	//nums[lo] = tempx
	return j
}
func sortq(nums[]int, lo int, hi int) {
	if lo >= hi {
		return
	}
	var  pivo  = partition(nums, lo, hi)
	sortq(nums, lo, pivo)
	sortq(nums, pivo+1, hi)
}
func sortArrayQuick(nums []int) []int {
	sortq(nums, 0, len(nums) - 1)
	return nums
}

// -------- 堆排序 -----------
func exch(nums []int, i int, j int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}

func less(nums []int, i int, j int) bool {
	if nums[i] < nums[j] {
		return true
	} else {
		return false
	}
}


func sink(nums []int, k int, N int) {
	/*
	@params nums 数组
	@params k 当前位置下标
	@params N 数组总长度
	 */
	for 2*k + 1 <= N {
		j := 2*k + 1
		if j + 1 <= N && nums[j] < nums[j+1] {
			j++
		}
		if !(nums[k] < nums[j]) {
			break
		}
		exch(nums, j, k)
		k = j
	}
}

func heapify(nums []int){
	// 数组 -> 堆有序
	var lenN int
	lenN = len(nums)
	for i := (lenN - 1) / 2; i >= 0; i-- {
		sink(nums, i, lenN -  1)
	}
}

func sortArray(nums []int) []int {
	// 初始化堆-堆有序
	var (
		i int
		lenN int
	)
	lenN = len(nums)
	heapify(nums)

	for i=lenN - 1; i >= 1; {
		exch(nums, 0, i)
		i--
		sink(nums, 0, i)
	}
	return nums
}