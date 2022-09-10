/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    cur3:=l1
    cur4:=l2
    
    add:=0
    for cur3!=nil && cur4!=nil {
        val:=cur3.Val+cur4.Val+add
        if val > 9 {
            temp:=val
            val=temp%10
            add=temp/10
            if add!=0 {
                switch {
                    case cur4.Next==nil:
                        newNode:=ListNode{Val:add}
                        cur4.Next=&newNode
                        add=0
                    case cur3.Next==nil:
                        newNode:=ListNode{Val:add}
                        cur3.Next=&newNode
                        add=0
                }
            }
        } else {
            switch {
                case cur4.Next==nil && cur3.Next!=nil:
                    newNode:=ListNode{}
                    cur4.Next=&newNode
                case cur3.Next==nil && cur4.Next!=nil:
                    newNode:=ListNode{}
                    cur3.Next=&newNode
            }
            add=0
        }
        cur4.Val=val
        cur3=cur3.Next
        cur4=cur4.Next
    }
    
    return l2
}   