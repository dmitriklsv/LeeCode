func merge(nums1 []int, m int, nums2 []int, n int)  {
    if n > 0 {
        length := m+n
        arr := make([]int,length)
        
        for i:=0;i<length;i++ {
            if i < m {
                arr[i]=nums1[i]
            } else {
                arr[i]=nums2[i%n]
            }
        }
        sort.Ints(arr)
        for i:=0;i<length;i++{
            nums1[i]=arr[i]
        }
    } 
}