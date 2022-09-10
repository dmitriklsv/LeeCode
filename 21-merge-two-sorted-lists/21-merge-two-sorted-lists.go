/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    if list1 == nil && list2 == nil {
        return nil
    }
    
    arr := []int{}
    cur1 := list1
    cur2 := list2
    
    for cur1 != nil {
        arr = append(arr, cur1.Val)
        cur1 = cur1.Next
    }
    
    for cur2 != nil {
        arr = append(arr, cur2.Val)
        cur2 = cur2.Next
    }

    for i := 0; i < len(arr) - 1; i++ {
        for j := 0; j < len(arr) - i - 1; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }    

    list3 := &ListNode{Val: arr[0]}
    cur3 := list3
    
    for k := 1; k < len(arr); k++ {
        newNode := ListNode{Val: arr[k]}
        for cur3.Next != nil {
            cur3 = cur3.Next
        }
        cur3.Next = &newNode
    }    
    
    return list3
}