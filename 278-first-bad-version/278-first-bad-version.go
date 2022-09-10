/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
    
    for i:=n; i>=1; i-- {
        if isBadVersion(i) && !isBadVersion(i-1){
            return i            
        }
    }
    return n
}