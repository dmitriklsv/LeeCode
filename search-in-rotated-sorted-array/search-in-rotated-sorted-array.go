func search(nums []int, target int) int {
    left:=0
    right:=len(nums)-1
    for left<=right {
        mid:=left+(right-left)/2
        switch {
            case nums[mid]>target:
                if target<nums[left] && nums[mid]>nums[right] {
                    left=mid+1
                } else {
                    right=mid-1
                }
            case nums[mid]<target: 
                if target>nums[right] && nums[mid]<nums[left] {
                    right=mid-1
                } else {
                    left=mid+1
                }
            default:
            return mid
        }
    }
    
    return -1
}

// Input: nums = [4,5,6,7,0,1,2], target = 0
// Output: 4