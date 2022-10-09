func countAndSay(n int) string {
    res:="1"
    return findSeq(n, res)
}

func findSeq(n int, res string) string {
    if n==1 {
        return res
    }

    num:=res[0]
    temp:=""
    c:='1'
    for i:=0;i<len(res);i++ {
        if i==len(res)-1 {
            temp+=string(c)+string(num)
        } else {
            if res[i]==res[i+1] {
                c++
            } else {
                temp+=string(c)+string(num)
                c='1'
                num=res[i+1]
            }
        }
    }
    
    return findSeq(n-1,temp)
}