func lengthOfLastWord(s string) int {
    chain,c:=false,0
    for i:=len(s)-1; i>=0; i-- {
        if s[i]!=32 {
            chain=true
            c++
        } else if chain {
            break
        } 
    }
    
    return c
}