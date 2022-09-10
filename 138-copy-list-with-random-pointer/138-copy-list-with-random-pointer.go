/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    if head==nil{
        return nil
    }
    copyHead:=*head
    cur:=head
    cur2:=&copyHead
    nodes:=make(map[*Node]*Node)
    
    for cur.Next!=nil{
        nodes[cur]=cur2
        cur=cur.Next
        temp:=*cur
        cur2.Next=&temp
        cur2=cur2.Next
    }
    nodes[cur]=cur2
    
    cur3:=&copyHead
    for cur3!=nil{
        if adrss,ok:=nodes[cur3.Random]; ok {
            cur3.Random=adrss
        }
        cur3=cur3.Next
    }    
    
    return &copyHead
}
//