func checkIfExist(arr []int) bool {
    for i:=0;i<len(arr);i++{
        for j:=i+1;j<len(arr);j++{
            if arr[i]==(arr[j])*2 || arr[j]==(arr[i])*2 {
                return true
            }
        }
    }
    return false
}