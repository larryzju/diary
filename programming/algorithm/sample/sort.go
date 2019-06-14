package sample

// MergeSortArray merge nums2 into nums1 (in place)as one sorted array
// - nums1 has enough space to hold additional elements from nums2
// - nums1 and nums2 are sorted already
func MergeSortArray(nums1, nums2 []int, len1, len2 int) {
	i := len1 - 1
	j := len2 - 1
	for t := i + j + 1; t >= 0; t-- {
		if i > 0 && j > 0 {
			if nums1[i] > nums2[j] {
				nums1[t] = nums1[i]
				i -= 1
			} else {
				nums1[t] = nums2[j]
				j -= 1
			}
		} else if i > 0 {
			nums1[t] = nums1[i]
			i -= 1
		} else if j > 0 {
			nums1[t] = nums2[j]
			j -= 1
		} else {
			break
		}
	}
}
