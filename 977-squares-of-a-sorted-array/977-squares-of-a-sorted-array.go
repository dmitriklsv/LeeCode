func sortedSquares(nums []int) []int {
    var res []int
    for _, nb := range nums {
        res = append(res,nb*nb)
    }
    
    for i:=0; i< len(res)-1; i++ {
        for j:=0; j<len(res)-1; j++ {
            if res[j] > res[j+1] {
                res[j],res[j+1] = res[j+1],res[j]
            }
        }
    }
        return res
}