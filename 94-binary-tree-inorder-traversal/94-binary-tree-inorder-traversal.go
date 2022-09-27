/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    res:=[]int{}
    traverse(root,&res)
    return res
}

func traverse(root *TreeNode, res *[]int) {
    if root == nil {
        return
    }
    traverse(root.Left,res)
    *res=append(*res,root.Val)
    traverse(root.Right,res)
}