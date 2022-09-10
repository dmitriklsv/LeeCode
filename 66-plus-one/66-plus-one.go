func plusOne(n []int) []int {
    temp:=1
    for i:=len(n)-1;i>=0;i--{
        if n[i]+temp>9{
            temp2:=n[i]
            n[i]+=temp-10
            temp=(temp+temp2)/10
        } else {
            n[i]+=temp
            temp=0
        }
    }
    // fmt.Println(n,temp)
    // res:=make([]int,len(n)+temp)
    if temp>0{
        res:=[]int{}
        res=append(res,temp)
        res=append(res,n...)
        return res
    }
    
    return n
}