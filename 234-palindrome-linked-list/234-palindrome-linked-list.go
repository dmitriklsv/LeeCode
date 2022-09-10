/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    if head.Next == nil {
        return true
    }
    var prev, next *ListNode
    slow := head
    fast := head
    firstHalf := *head
    curr1 := &firstHalf
    
    for slow != nil && fast != nil && fast.Next != nil && fast.Next.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        curr1 = curr1.Next
    }
    secondHalf := *curr1.Next
    curr1.Next = nil
    curr2 := &secondHalf
    
    for curr2 != nil {
        next = curr2.Next
        curr2.Next = prev
        prev = curr2
        curr2 = next
    }
    curr3 := prev
    curr4 := &firstHalf    
    
    for curr3 != nil && curr4 != nil {
        if curr3.Val != curr4.Val {
            return false
        }
        curr3 = curr3.Next
        curr4 = curr4.Next
    }
    
    return true
}