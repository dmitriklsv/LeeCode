func isPalindrome(x int) bool {
	strX := strconv.Itoa(x)
	if len(strX) == 1 {
		return true
	} else if len(strX) == 2 {
		if strX[0] == strX[1] {
			return true
		} else {
			return false
		}
	} else if len(strX)%2 != 0 {
		var firstHalf, secondHalf []string
		for i := 0; i < (len(strX)-1)/2; i++ {
			firstHalf = append(firstHalf, string(strX[i]))
		}
		for i := len(strX) - 1; i > len(strX)/2; i-- {
			secondHalf = append(secondHalf, string(strX[i]))
		}
		
		var flag bool
		for i := 0; i <= len(firstHalf)-1; i++ {
			if firstHalf[i] == secondHalf[i] {
				flag = true
			} else {
				return false
			}
		}

		if flag {
			return true
		}
		return false

	} else {
		var firstHalf, secondHalf []string
		for i := 0; i < len(strX)/2; i++ {
			firstHalf = append(firstHalf, string(strX[i]))
		}
		for i := len(strX) - 1; i >= len(strX)/2; i-- {
			secondHalf = append(secondHalf, string(strX[i]))
		}
		
		var flag bool
		for i := 0; i <= len(firstHalf)-1; i++ {
			if firstHalf[i] == secondHalf[i] {
				flag = true
			} else {
				return false
			}
		}

		if flag {
			return true
		}
		return false
	}
}
