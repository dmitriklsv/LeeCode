func numIslands(grid [][]byte) int {
    var c int
    for i:=0;i<len(grid);i++{
        for j:=0;j<len(grid[i]);j++ {
            if grid[i][j] == '1' {
                dfs(&grid,i,j)
                c++
            }
        }
    }
    
    return c
}

func dfs(grid *[][]byte,i,j int) {
    if  i < 0 || i > len(*grid)-1 || j > len((*grid)[i])-1 || j < 0 || (*grid)[i][j] != '1' {
        return
    }
    
    (*grid)[i][j] = 'x'
    
    dfs(grid,i,j+1)
    dfs(grid,i+1,j)
    dfs(grid,i,j-1)
    dfs(grid,i-1,j)
}