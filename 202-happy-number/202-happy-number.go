func isHappy(n int) bool {
    table:=make(map[int]struct{})
    for n != 1 {
        if _,ok := table[n]; ok {
            return false
        }
        table[n]=struct{}{}
        num:=0
        for n>0 {
            num+=(n%10)*(n%10)
            n/=10
        }
        n=num
    }
    
    return true
}