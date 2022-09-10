func minSubArrayLen(target int, nums []int) int {
    if nums[len(nums)-1] >= target {
        return 1
    }
    i,j,temp,count,found:=0,1,nums[0],len(nums),false
    for i<len(nums)-1 {
        if j==len(nums) {
            if !found {
                return 0
            } else {
                return count
            }
        }
        if temp >= target {
            return 1
        }
        if temp+nums[j]<target {
            temp+=nums[j]
            j++
        } else if temp+nums[j]>=target {
            if j-i+1<=count{
                found=true
                count=j-i+1
            }
            i++
            j=i+1
            temp=nums[i]
        }
    }
    
    if found {
        return count
    }
    return 0
}
//Input: target = 7, nums = [2,3,1,2,4,3]
// Output: 2