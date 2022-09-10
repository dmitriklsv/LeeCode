func validMountainArray(arr []int) bool {
    if len(arr) < 3 {
        return false
    }
    
    up,down:=false,false
    countUp,countDown:=0,0
    for i:=0;i<len(arr)-1;i++{
        if arr[i+1]==arr[i] || arr[0]>arr[1] {
            return false
        }
        
        if arr[i+1]>arr[i] {
            if down {
                countDown++
                down=false
            }
            up=true
        } else if arr[i+1]<arr[i] {
            if up {
                countUp++
                up=false
            }
            down=true
        }
        
        if i+1==len(arr)-1 {
            if down {
                countDown++
            } else if up {
                countUp++
            }
        } 
    }
    
    if countUp!=1 || countDown!=1 {
        return false
    }
    return true
}