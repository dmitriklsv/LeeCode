func minCost(colors string, neededTime []int) int {
    res, max := 0, neededTime[0]
    for i:=0;i<len(colors)-1;i++{
        if colors[i]==colors[i+1] {
            if max<neededTime[i+1] {
                res+=max
                max=neededTime[i+1]
            } else {
                res+=neededTime[i+1]
            }
        } else {
            max=neededTime[i+1]
        }
    }
    return res
}