func pivotIndex(nums []int) int {
    for i:=0;i<len(nums);i++{
        left,right:=i,i
        leftSum,rightSum:=0,0
        for left>0 {
            left--
            leftSum += nums[left]
        }
        
        for right<len(nums)-1 {
            right++
            rightSum += nums[right]
        }
        
        if leftSum-rightSum==0 {
            return i
        }
    }
    
    return -1
}