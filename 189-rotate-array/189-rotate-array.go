func rotate(nums []int, k int) {
	if len(nums) == 1 {
		return
	}
	rot := k % len(nums)
	temp := nums[len(nums)-rot:]
	temp = append(temp, nums[:len(nums)-rot]...)
	for i := range temp {
		nums[i] = temp[i]
	}
}
