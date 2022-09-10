
func lengthOfLongestSubstring(s string) int {
	i, count := 0, 0
	table := map[byte]bool{}
	unique := true

	for i < len(s) {
		_, ok := table[s[i]]
		if !ok {
			table[s[i]] = true
		} else {
			unique = false
            if len(table) >= count {
                count = len(table)
			}
			i -= len(table)
			for k := range table {
				delete(table, k)
			}
		}
		i++
	}

	if unique {
		return len(table)
	} else if len(table) >= count {
		return len(table)
	} else {
		return count
	}
}
