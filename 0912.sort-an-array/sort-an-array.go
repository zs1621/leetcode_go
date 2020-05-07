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

// 归并排序入口
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

func sort(nums []int, low int, high int) {
	var middle int
	if high-low <= InsertSortThreshold {
		sortInsert(nums, low, high)
		return
	}
	if low < high {
		middle = low + (high-low)/2
		sort(nums, low, middle)
		sort(nums, middle+1, high)
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

func sortArray(nums []int) []int {
	// 归并排序入口
	tempArray = make([]int, len(nums))
	sort(nums, 0, len(nums)-1)
	return nums
}
