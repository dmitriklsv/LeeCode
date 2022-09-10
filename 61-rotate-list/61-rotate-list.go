/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    if head ==nil{
        return nil
    }
    if head.Next==nil{
        return head
    }
    
    temp:=head
    cur:=head
    some:=head
    var prev *ListNode
    length:=0
    for some!=nil{
        some=some.Next
        length++
    }
    k=k%length

    for k>0{
        for cur.Next!=nil{
            prev=cur
            cur=cur.Next
        }
        head=cur
        head.Next=temp
        prev.Next=nil
        prev=cur
        temp=cur
        k--
    }
    
    return head
}