func dominantIndex(nums []int) int {
    max,pos:=nums[0],0
    for i:=1;i<len(nums);i++{
        if nums[i]>max{
            max=nums[i]
            pos=i
        }
    }

    for i, v := range nums {
        if v*2 > max && i != pos {
            return -1
        }
    }
    
    return pos
}