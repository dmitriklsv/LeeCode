func sortedSquares(nums []int) []int {
    arr := make([]int,len(nums))
    for i, nb := range nums {
        arr[i]=nb*nb
    }
    sort.Ints(arr)
    return arr
}