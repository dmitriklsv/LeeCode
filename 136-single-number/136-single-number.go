func singleNumber(nums []int) int {
    table := make(map[int]int)
    for _,v := range nums {
        table[v]++
    }
    for k,v := range table {
        if v == 1 {
            return k
        }
    }
    return -1
}