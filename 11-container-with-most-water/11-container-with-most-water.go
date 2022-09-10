func maxArea(height []int) int {
    length:=len(height)-1
    i,j:=0,length
    mostWater:=0
    for i<j {
        minElem:=height[j]
        if height[i]<height[j] {
            minElem=height[i]
            if minElem*(j-i)>mostWater {
                mostWater=minElem*(j-i)
            }
            i++
            continue
        }
        if minElem*(j-i)>mostWater {
            mostWater=minElem*(j-i)
        }
        j--
    }

    
    return mostWater
}