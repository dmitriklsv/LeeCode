func duplicateZeros(arr []int)  {
    temp := make([]int,len(arr))
    length := 0
    for i:=0;i<len(arr);i++ {
        if length >= len(arr) {
            break
        }
        temp[length] = arr[i]
        length++
        if arr[i]==0 {
            length++
        }
    }
    // fmt.Println(temp)
    for i:=0;i<len(arr);i++{ 
        arr[i]=temp[i]
    }
}