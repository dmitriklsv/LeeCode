func searchInsert(nums []int, target int) int {
    var match bool
    var res,expect int
    for i:=0; i < len(nums); i++ {
        if target > nums[i]{
            expect = i + 1            
        } 
        if nums[i] == target {
            match = true
            res = i
        }
    }
    
    if match {
        return res
    } else {
        return expect
    }
}

