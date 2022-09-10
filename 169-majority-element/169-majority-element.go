func majorityElement(nums []int) int {
    table:=make(map[int]int)
    for _,v:=range nums {
        table[v]++
    }
    maxV:=0
    maxK:=0
    for k,v:=range table{
        if v>maxV{
            maxV=v
            maxK=k
        }
    }
    return maxK
}