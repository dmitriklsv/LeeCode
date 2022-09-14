func threeSumClosest(nums []int, target int) int {
    sort.Ints(nums)
    closest:=nums[0]+nums[1]+nums[2]
    for i:=0;i<len(nums)-2;i++{
        // fmt.Println("hi")
        if i!=0 && nums[i]==nums[i-1] {
            continue
        }
        j:=i+1
        k:=len(nums)-1
        for j<k {
            sum:=nums[i]+nums[j]+nums[k]
            if sum==target {
                return target
            }
            if sum<target{
                j++
                for j<len(nums)-1 && nums[j]==nums[j-1] {
                    j++
                }
            } else if sum>target {
                k--
            }
            if math.Abs(float64(target-sum))<math.Abs(float64(target-closest)) {
                closest=sum
            }
        }
    }
    
    return closest
}