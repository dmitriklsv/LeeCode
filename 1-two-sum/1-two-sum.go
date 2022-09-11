func twoSum(nums []int, t int) []int {
    table:=make(map[int]int)
    for i:=0;i<len(nums);i++{
        rN:=t-nums[i]
        if j,ok:=table[rN];ok{
            return []int{j,i}
        }
        table[nums[i]]=i
    }
    
    return nil
}