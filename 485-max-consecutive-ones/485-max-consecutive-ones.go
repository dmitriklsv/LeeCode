func findMaxConsecutiveOnes(nums []int) int {
    count, temp := 0, 0
    consec := false
    for i := 0; i < len(nums); i++ {
        if nums[i] == 1 {
            temp++
            consec = true
            if i==len(nums)-1 && consec && temp > count {
                return temp
            }
        } else if consec {
            consec = false
            if temp > count {
                count = temp
            }
            temp = 0
        }
    }
    return count
}