func numIslands(grid [][]byte) int {
    var c int
    var visit [][]int
   
    for i:=0;i<len(grid);i++ {
        for j:=0;j<len(grid[i]);j++ {
            if grid[i][j] == '1' {
                visit = append(visit,[]int{i,j})
                bfs(&grid,i,j,&visit)
                c++
            }
        }
    }
    
    return c
}

func bfs(grid *[][]byte,i,j int,visit *[][]int) {
    if j<len((*grid)[i])-1 && (*grid)[i][j+1]=='1'{
        *visit=append(*visit,[]int{i,j+1})
        (*grid)[i][j+1]='x'
    }
    if i<len(*grid)-1 && (*grid)[i+1][j]=='1'{
        *visit=append(*visit,[]int{i+1,j})
        (*grid)[i+1][j]='x'
    }
    if j>0 && (*grid)[i][j-1]=='1' {
        *visit=append(*visit,[]int{i,j-1})
        (*grid)[i][j-1]='x'
    }
    if i>0 && (*grid)[i-1][j]=='1' {
        *visit=append(*visit,[]int{i-1,j})
        (*grid)[i-1][j]='x'
    }
    
    *visit=(*visit)[1:]
    (*grid)[i][j]='x'
    
    if len(*visit) != 0 {
        bfs(grid,(*visit)[0][0],(*visit)[0][1],visit)
    }
}