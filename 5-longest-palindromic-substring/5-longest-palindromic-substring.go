func longestPalindrome(s string) string {
    res := s[0:1]
    for i := range s {
        res = checkStep(i, i, s, res)
        res = checkStep(i, i+1, s, res)
    }
    
    return res
}

func checkStep(start, end int, s, res string) string {
    for start >= 0 && end < len(s) && s[start] == s[end] {
        start--
        end++
    }
    
    if end - start - 1 > len(res) {
        return s[start+1:end]
    }
    
    return res
}