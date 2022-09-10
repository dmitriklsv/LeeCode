func longestCommonPrefix(strs []string) string {
    count,res:=0,""
    
    for len(strs) > 1 {
        found:=0
        for i:=0;i<len(strs)-1;i++{
            if count>=len(strs[i]) || count>=len(strs[i+1]) {
                return res
            }
            if strs[i][count]==strs[i+1][count]{
                found++
            } else {
                found=0
            }
            if found==len(strs)-1 {
                res+=string(strs[i][count])
            } else if found ==0{
                return res
            } 
        }
        count++
    }
    
    return strs[0]    
}