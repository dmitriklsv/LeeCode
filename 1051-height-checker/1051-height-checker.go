func heightChecker(n []int) int {
    arr:=make([]int,len(n))
    for i,v:=range n {
        arr[i]=v
    }
    for i:=0;i<len(n)-1;i++{
        for j:=0;j<len(n)-1-i;j++{
            if n[j] > n[j+1] {
                n[j],n[j+1]=n[j+1],n[j]
            }
        }
    }    
    res:=0
    for i:=range n {
        if n[i]!=arr[i] {
          res++  
        }
    }
    
    // fmt.Println(arr,"\n",n)
    return res
}