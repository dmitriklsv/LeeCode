func isValid(s string) bool {   
    stack:=[]byte{}
    for i:=0;i<len(s);i++{
        if s[i]=='('||s[i]=='{'||s[i]=='[' {
            stack=append(stack,s[i])
        } else {
            if len(stack)==0 {
                return false
            }
            last:=stack[len(stack)-1]
            stack=stack[:len(stack)-1]
            if (s[i]==')' && last!='(') || (s[i]=='}' && last!='{') || (s[i]==']' && last!='[') {
                return false
            }
        }
    }
    
    return len(stack)==0
}