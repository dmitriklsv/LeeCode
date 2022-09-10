func search(nums []int, target int) int {
    left := 0
    right := len(nums)-1
    
    for left <= right {
        mid := left + (right - left)/2

        switch {
        case target == nums[mid]: return mid
        case target < nums[mid]: right = mid - 1
        default: left = mid + 1
        }
    }
    
    return -1
}