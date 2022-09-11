func searchMatrix(matrix [][]int, target int) bool {
    for i:=0;i<len(matrix);i++{
        left:=0
        right:=len(matrix[i])-1
        for left<=right{
            mid:=left+(right-left)/2
            switch {
                case matrix[i][mid]==target: return true
                case matrix[i][mid]<target: left = mid + 1
                default: right = mid - 1
            }
        }
    }
    
    return false 
}