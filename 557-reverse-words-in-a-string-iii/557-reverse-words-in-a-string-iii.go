func reverseWords(s string) string {
var isChar bool
	var temp, res string
	for i, c := range s {
		if i == len(s)-1 {
			res += string(c) + temp
		}
		if c != 32 {
			isChar = true
			temp = string(c) + temp
		} else if isChar {
			isChar = false
			res += temp + string(c)
			temp = ""
		}
	}
	return res
}