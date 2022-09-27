type MyHashSet struct {
    table map[int]bool
}


func Constructor() MyHashSet {
    return MyHashSet{make(map[int]bool)}
}


func (this *MyHashSet) Add(key int)  {
    if !this.Contains(key) {
        this.table[key]=true
    }
}


func (this *MyHashSet) Remove(key int)  {
    if this.Contains(key) {
        delete(this.table,key)
    }
}


func (this *MyHashSet) Contains(key int) bool {
    if this.table[key] {
        return true
    }
    return false
}


/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */