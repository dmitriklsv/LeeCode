func findDiagonalOrder(mat [][]int) []int {
    table:=make(map[int][]int)
    for i:=0;i<len(mat);i++{
        for j:=0;j<len(mat[i]);j++{
            table[i+j]=append(table[i+j],mat[i][j])
        }
    }
    
    res:=make([]int,len(mat)*len(mat[0]))
    k:=0
    for i:=0;i<len(table);i++{
        if i%2==0{
            for j:=len(table[i])-1;j>=0;j--{
                res[k]=table[i][j]
                k++
            }
        } else {
            for j:=0;j<len(table[i]);j++{
                res[k]=table[i][j]
                k++
            }
        }
    }
    
    return res
}