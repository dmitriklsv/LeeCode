func findKthPositive(arr []int, k int) int {
  table := make(map[int]bool, 1000)
	for n := 1; n <= 10000; n++ {
		_, ok := table[n]
		if !ok {
			table[n] = true
		}
	}

	keys := make([]int, 0, len(table))
	for k := range table {
		keys = append(keys, k)
	}
	sort.Ints(keys)
    
    for _, nb := range arr {
		_, ok := table[nb]
		if ok {
			delete(table, nb)
		}
	}
    
    count,res := 1,0
	for _, c := range keys {
		_, ok := table[c]
		if ok {
			if count == k {
                res = c
                break
			}
			count++
		}
	}
    
    return res
}

