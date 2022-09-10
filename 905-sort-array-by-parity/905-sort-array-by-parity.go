func sortArrayByParity(nums []int) []int {
    count:=0
    for i:=0;i<len(nums);i++{
        if nums[i]%2==0 {
            for j:=i;j>count;j-- {
                nums[j-1],nums[j]=nums[j],nums[j-1]
            }
            count++
        }
    }
    return nums
}