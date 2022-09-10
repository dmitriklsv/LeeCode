func reverseWords(s string) string {
    arr:=[]string{}
    temp:=""
    found:=false
    for _,ch := range s {
        if ch!=' ' {
            if !found {
                found=true
            }
            temp+=string(ch)
        } else if found {
            arr=append(arr,temp)
            temp=""
            found=false
        }
    }
    
    if found {
        arr=append(arr,temp)
    }

    res:=""
    for i:=len(arr)-1;i>=0;i--{
        res+=arr[i] + " "
    }
    
    return res[:len(res)-1]
}
