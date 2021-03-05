// Median of Two Sorted Arrays
// 
// [Hard] [AC:27.1% 492.5K of 1.8M] [filetype:golang]
// 
// There are two sorted arrays nums1 and nums2 of size m and n
// respectively.
// 
// Find the median of the two sorted arrays. The overall run time
// complexity should be O(log (m+n)).
// 
// You may assume nums1 and nums2 cannot be both empty.
// 
// Example 1:
// 
// nums1 = [1, 3]
// 
// nums2 = [2]
// 
// The median is 2.0
// 
// Example 2:
// 
// nums1 = [1, 2]
// 
// nums2 = [3, 4]
// 
// The median is (2 + 3)/2 = 2.5
// 
// [End of Description]
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var(
		index,
		index1,
		index2 int,
	)
	len1 := len(nums1)
	len2 := len(nums2)
	for{
		if index > (len1 + len2)/2{
			break
		}
		if nums1[index1] > nums2[index2] && index2 < len2{
			index2 ++
		}else{
			index1 ++
		}
		index ++
	}
}
