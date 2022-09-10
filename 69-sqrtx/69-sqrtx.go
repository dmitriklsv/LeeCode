func mySqrt(x int) int {
	left := 0
	right := x
	for left <= right {
		mid := left + (right-left)/2
		switch {
		case mid*mid > x:
			right = mid - 1
		case mid*mid < x:
			left = mid + 1
		default:
			return mid
		}
	}
	return right
}