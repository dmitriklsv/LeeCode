/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
    var prev *ListNode
    curr := head
    iter := false
    
    for curr != nil {
        if curr.Val == val {
            if !iter {
                head = head.Next
                curr = curr.Next
                continue
            } else {
                prev.Next = curr.Next
                curr = curr.Next
                continue
                // iter = false
            }
        }
        
        prev = curr
        curr = curr.Next
        iter = true
    }
    return head
    
    // curr := head
    // for curr!=nil {
    //     head = head.Next
    //     curr=curr.Next
    // }
    // return head
}