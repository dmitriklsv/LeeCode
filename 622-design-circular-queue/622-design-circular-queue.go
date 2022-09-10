type MyCircularQueue struct {
    data []int
    head,tail,size int
}


func Constructor(k int) MyCircularQueue {
    return MyCircularQueue{make([]int,k),-1,-1,k}
}


func (this *MyCircularQueue) EnQueue(value int) bool {
    if this.IsFull() {
        return false
    }
    if this.IsEmpty() {
        this.head=0
    }
    this.tail=(this.tail+1)%this.size
    this.data[this.tail]=value
    return true
}


func (this *MyCircularQueue) DeQueue() bool {
    if this.IsEmpty() {
        return false
    }
    if this.head==this.tail {
        this.head=-1
        this.tail=-1
        return true
    }
    this.head=(this.head+1)%this.size
    return true
}


func (this *MyCircularQueue) Front() int {
    if this.IsEmpty() {
        return -1
    }
    return this.data[this.head]
}


func (this *MyCircularQueue) Rear() int {
    if this.IsEmpty() {
        return -1
    }
    return this.data[this.tail]
}


func (this *MyCircularQueue) IsEmpty() bool {
    return this.head==-1
}


func (this *MyCircularQueue) IsFull() bool {
    return (this.tail+1)%this.size==this.head
}


/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */