func intersection(nums1 []int, nums2 []int) []int {
    table:=make(map[int]bool)
    res:=[]int{}
    for _,v:=range nums1 {
        table[v]=true
    }
    for _,v:=range nums2 {
        if _,ok:=table[v]; ok {
            res=append(res,v)
            delete(table,v)
        }
    }
    return res
}