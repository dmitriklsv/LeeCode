func reverse(x int) int {
    
    negative:=false
    if x<0 {
        x*=-1
        negative=true
    }
    res:=0
    for x>0 {
        res*=10
        res+=x%10
        x/=10
    }
    
    if res<math.MinInt32 || res>math.MaxInt32 {
        return 0
    }
    
    if negative {
        return -res
    }
    return res
}