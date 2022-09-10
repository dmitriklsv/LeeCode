func removeDuplicates(nums []int) int {
    temp,count:=nums[0],1
    arr:=[]int{}
    arr=append(arr,temp)
    for i:=1;i<len(nums);i++{
        if nums[i]!=temp {
            temp = nums[i]
            arr=append(arr,temp)
            count++
        }
    }
    
    for i:=0;i<count;i++ {
        nums[i]=arr[i]
    }
    
    return count
}