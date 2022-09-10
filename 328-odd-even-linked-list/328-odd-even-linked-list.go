/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    
    var oddnext, evennext *ListNode
    odd := head
    even := head.Next
    evenSt := head.Next
    
    for odd != nil && even != nil && odd.Next.Next != nil {
        oddnext = odd.Next.Next
        evennext = even.Next.Next
        odd.Next = odd.Next.Next
        even.Next = even.Next.Next
        odd = oddnext
        even = evennext
    }
    odd.Next = evenSt
    
    return head
}