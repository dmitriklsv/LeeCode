func containsDuplicate(nums []int) bool {
    table:=make(map[int]bool)
    for _,v:=range nums {
        if _,ok:=table[v];!ok {
            table[v]=true
        } else {
            return true
        }
    }
    return false
}