func removeElement(nums []int, val int) int {
    count := 0
    for _, nb := range nums {
        if nb != val {
            count++
        }
    }
    
    iter := 0
    for i := 0; i < len(nums) - 1; i++ {
        if iter > count {
            break
        }
        if nums[i] == val {
            iter++
            // iter = 0
            j := i
            for j < len(nums) - 1 {
                nums[j], nums[j+1] = nums[j+1], nums[j]
                j++
            }
            i--
        }
        
    }
    
    
    return count
}