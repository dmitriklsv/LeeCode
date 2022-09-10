func rotate(nums []int, k int)  {
    if len(nums) == 1 {
        return
    } else  {
        rot := k%len(nums)
    var temp []int
	for i := len(nums) - rot; i < len(nums); i++ {
		temp = append(temp, nums[i])
	}

	for j := 0; j < len(nums)-rot; j++ {
		temp = append(temp, nums[j])
	}

	for i := range nums {
		nums[i] = temp[i]
	}    
    }   
    }
	
