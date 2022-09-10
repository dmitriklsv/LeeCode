func convert(s string, numRows int) string {
    if numRows==1 {
        return s
    }
    
    arr:=make([]string,numRows)
    i,j:=0,0
    for i<len(s) {
        arr[j]+=string(s[i])
        j++ 
        i++
        if j==len(arr)-1 {
            for j>0 && i<len(s) {
                arr[j]+=string(s[i])
                j--
                i++
            }
        }
    }
    
    res:=""
    for _,v:=range arr {
        res+=v
    }
    
    return res
}
