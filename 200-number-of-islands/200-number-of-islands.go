func numIslands(grid [][]byte) int {
    numIslands:=0
    for i:=0;i<len(grid);i++{
        for j:=0;j<len(grid[i]);j++{
            if grid[i][j]=='1' {
                numIslands++
                dfs(&grid,i,j)
            }
        }
    }
    
    return numIslands
}

func dfs(grid *[][]byte, row, col int) {
    if row<0 || col<0 || row>=len(*grid) || col>=len((*grid)[row]) || (*grid)[row][col]=='0'{
        return
    }
    
    (*grid)[row][col]='0'
    dfs(grid,row-1,col)
    dfs(grid,row,col-1)
    dfs(grid,row+1,col)
    dfs(grid,row,col+1)
}
