func addBinary(a string, b string) string {
    length:=len(b)
    if len(a)>len(b) {
        length = len(a)
        temp:=b
        for i:=length; i>len(b);i-- {
            temp="0"+temp
        }
        b=temp
    } else if len(b)>len(a) {
        temp:=a
        for i:=length; i>len(a);i--{
           temp ="0"+temp
        }
        a=temp
    }
    
    add, res :=0, ""
    for i:=length-1;i>=0;i--{
        temp:=0
        temp=int(a[i]-48)+int(b[i]-48)+add
        if temp==2{
            temp=0
            add=1
        } else if temp==3{
            temp=1
            add=1
        } else {
            add=0
        }
        
        res = string(rune(temp+48)) + res
    }
    
    if add==1{
        res=string(rune(add+48))+res
    }
    
    return res
}
