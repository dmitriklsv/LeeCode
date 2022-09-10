func findNumbers(nums []int) int {
    count:=0
    for _, num := range nums {
        length := 0
        for num > 0 {
            length++
            num/=10
        }
        if length%2==0 {
            count++
        }
    }
    return count
}