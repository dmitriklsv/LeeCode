func findDisappearedNumbers(nums []int) []int {
    length:=len(nums)
    arr:=make([]int,length)
    for i:=0;i<length;i++{
        arr[i]=i+1
    }
    
    res:=[]int{}
    for i:=0;i<length;i++{
        found:=false
        for j:=0;j<length;j++{
            if arr[i]==nums[j] {
                found=true
                break
            }
        }
        if !found{
            res=append(res,arr[i])
        }
    }
    return res
}