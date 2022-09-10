/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Prev *Node
 *     Next *Node
 *     Child *Node
 * }
 */

func flatten(root *Node) *Node {
    cur:=root
    for cur!=nil{
        if cur.Child!=nil{
            temp:=cur.Child
            next:=cur.Next
            cur.Next=cur.Child
            cur.Child.Prev=cur
            for temp.Next!=nil{
                temp=temp.Next
            }
            temp.Next=next
            if next!=nil{
                next.Prev=temp
            }
            cur.Child=nil
        }
        cur=cur.Next
    }
    
    return root
}