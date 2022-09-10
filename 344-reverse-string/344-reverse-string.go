func reverseString(s []byte)  {
    printReverse(s,0,len(s)-1)
}

func printReverse(s []byte,i,j int) {
    if i>=j {
        return
    }
    s[i],s[j]=s[j],s[i]
    printReverse(s,i+1,j-1)
}
