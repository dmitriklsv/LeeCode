func climbStairs(n int) int {
    f0, f1 := 0, 1
    fn:=f0+f1
    
    for n > 0 {
        // fmt.Println(fn,f0,f1)
        f1 = f0
        f0 = fn
        fn = f0 + f1
        n--
    }
    
    return fn
}