func interpret(command string) string {
	res := ""
	for i := 0; i < len(command); i++ {
		if i != len(command)-1 && command[i] == '(' && command[i+1] == ')' {
			res += "o"
		} else if command[i] == ')' && i-1 >= 0 && i-2 >= 0 && i-3 >= 0 && command[i-1] == 'l' && command[i-2] == 'a' && command[i-3] == '(' {
			res += "al"
		} else if command[i] == 'G' {
			res += "G"
		}
	}
	return res
}
