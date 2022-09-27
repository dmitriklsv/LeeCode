type MyHashMap struct {
    table map[int]int
}


func Constructor() MyHashMap {
    return MyHashMap{make(map[int]int)}
}


func (this *MyHashMap) Put(key int, value int)  {
    this.table[key]=value
}


func (this *MyHashMap) Get(key int) int {
    if _,ok:=this.table[key];ok {
        return this.table[key]
    }
    return -1
}


func (this *MyHashMap) Remove(key int)  {
    delete(this.table,key)
}


/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */