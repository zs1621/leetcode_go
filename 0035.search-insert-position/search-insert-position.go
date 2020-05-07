package basic

// BSearch nums, target
func searchInsert(nums []int, target int) int {
	var (
		middle int
		i      int
		e      int
	)
	i = 0
	e = len(nums)
	for i < e {
		middle = i + (e-i)/2
		if nums[middle] < target {
			i = middle + 1
		} else {
			e = middle
		}
	}
	return i
}
