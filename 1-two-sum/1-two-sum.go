func twoSum(nums []int, target int) []int {
	var res []int
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if j == i {
				continue
			} else if nums[i]+nums[j] == target {
				res = append(res, i, j)
			}
		}
	}
	return res
}