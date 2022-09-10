func replaceElements(arr []int) []int {
    temp:=arr[len(arr)-1]
    temp2:=0
    for i:=len(arr)-1;i>0;i--{
        if temp<arr[i-1] {
            temp2 = temp
            temp = arr[i-1]
            arr[i-1]= temp2
        } else if temp>arr[i-1] {
            arr[i-1]=temp
        }
    }
    arr[len(arr)-1]=-1
    return arr
}