func isHappy(n int) bool {
    table:=make(map[int]bool)
    for n != 1 {
        if table[n] {
            return false
        }
        table[n]=true
        num:=0
        for n>0 {
            num+=(n%10)*(n%10)
            n/=10
        }
        n=num
    }
    
    return true
}