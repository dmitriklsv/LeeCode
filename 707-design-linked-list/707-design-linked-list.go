type MyLinkedList struct {
    head *node
    length int
}

type node struct {
    val int
    next *node
}


func Constructor() MyLinkedList {
    return MyLinkedList{}
}


func (this *MyLinkedList) Get(index int) int {
    if index >= this.length || index < 0 {
        return -1
    }
    
        switch index {
	case 0:
		return this.head.val
	default:
		count := 0
		curr := this.head

		for count != index {
			count++
			curr = curr.next
		}
		return curr.val
	}
}


func (this *MyLinkedList) AddAtHead(val int)  {
    newNode := node{val: val}
    second := this.head
    this.head = &newNode
    this.head.next = second
    this.length++
}


func (this *MyLinkedList) AddAtTail(val int)  {
    newNode := node{val: val}
    
    switch this.length {
        case 0:
            this.head = &newNode
        default:
            curr := this.head
            
            for curr.next != nil {
                curr=curr.next
            }
            curr.next = &newNode
    }
    this.length++
}


func (this *MyLinkedList) AddAtIndex(index int, val int)  {
    if index > this.length || index < 0 {
        return
    }
    
    switch index {
        case 0:
        this.AddAtHead(val)
    default:
        count:=0
        prev:=this.head
        curr:=this.head
        newNode:=node{val:val}
        
        for count!=index{
            count++
            prev=curr
            curr=curr.next
        }
        prev.next = &newNode
        newNode.next=curr
    }
    this.length++
}


func (this *MyLinkedList) DeleteAtIndex(index int)  {
    if index < 0 || index >= this.length {
        return
    }
    
    switch index {
        case 0:
            this.head = this.head.next
    default:
        count:=0
        prev:=this.head
        curr:=this.head
        
        for count!=index{
            count++
            prev=curr
            curr=curr.next
        }
        prev.next=curr.next
        curr=nil
    }
    this.length--
}


/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */