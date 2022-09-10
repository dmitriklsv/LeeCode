func romanToInt(s string) int {
    // symbols:="IVXLCDM"
    res:=0
    for i:=0;i<len(s);i++{
        switch s[i] {
        case 'I':
            if i+1<len(s) {
                if s[i+1]=='V' {
                    res+=4
                    i++
                    continue
                } else if s[i+1]=='X' {
                    res+=9
                    i++
                    continue
                }    
            } 
            res+=1
            
        case 'V':res+=5
        case 'X':
            if i+1<len(s) {
                if s[i+1]=='C' {
                    res+=90
                    i++
                    continue
                } else if s[i+1]=='L' {
                    res+=40
                    i++
                    continue
                } 
            }
            res+=10
            
        case 'L':res+=50
        case 'C':
            if i+1<len(s) {
                if s[i+1]=='M' {
                    res+=900
                    i++
                    continue
                } else if s[i+1]=='D' {
                    res+=400
                    i++
                    continue
                }
            }
            res+=100
            
        case 'D':res+=500
        case 'M':res+=1000
        }
    }
    
    return res
}