package shuffle_array

func shuffle(nums []int, n int) []int {
	numsLen := len(nums)
	ret := make([]int, numsLen)
	for idx := 0; idx < n; idx++ {
		ret[2*idx] = nums[idx]
		ret[2*idx+1] = nums[idx+n]
	}
	return ret
}
