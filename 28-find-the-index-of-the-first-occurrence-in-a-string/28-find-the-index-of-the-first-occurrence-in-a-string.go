func strStr(haystack string, needle string) int {
    count,pos,found:=0,-1,false
    for i:=0;i<len(haystack);i++{
        if haystack[i]==needle[count] {
            count++
        } else {
            i=i-count
            count=0
        }
        if count==len(needle) {
            temp:=i-count+1
            count=0
            if temp > pos && !found {
                pos = temp
                found=true
            }
        }
    }
    
    return pos
}