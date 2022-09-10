func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    arr := make([]int, len(nums1)+len(nums2))
	i := 0
	for _, v := range nums1 {
		arr[i] = v
		i++
	}
	for _, v := range nums2 {
		arr[i] = v
		i++
	}
	sort.Ints(arr)

	res := 0.0
	if len(arr)%2 != 0 {
		res = float64(arr[(len(arr)-1)/2])
	} else {
		var n1, n2 float64
		n1 = float64(arr[len(arr)/2-1])
		n2 = float64(arr[len(arr)/2])
		res = (n1 + n2) / 2
	}
	return res
}