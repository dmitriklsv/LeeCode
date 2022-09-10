/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    first := head
    count := 1
    
    for first.Next != nil {
        first = first.Next
        count++
    }
    
    second := head
    prev := head
    nodePos := count - n
    
    if nodePos == 0 {
        head = head.Next
        return head
    }
        // fmt.Println("yo")
    
    for nodePos > 0 {
        prev = second
        second = second.Next
        nodePos--
    }
    prev.Next = prev.Next.Next
    // second.Next = second.Next.Next
    
    return head
}