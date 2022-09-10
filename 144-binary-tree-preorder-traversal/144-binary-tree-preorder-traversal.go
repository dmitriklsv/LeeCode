/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
    res:=[]int{}
    for root!=nil{
        res=append(res,root.Val)
        
        if root.Left==nil{
            root=root.Right
            continue
        }
        
        if root.Right!=nil{
            cur:=root.Left
            for cur.Right!=nil{
                cur=cur.Right
            }
            cur.Right=root.Right
        }
        
        root=root.Left
    }        
    
    return res
}