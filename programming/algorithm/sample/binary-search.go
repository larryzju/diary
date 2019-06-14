package sample

// Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.
// Return the index of the v, or return -1 if it is not found in the array
func SearchInRotatedSortedArray(array []int, v int) int {
	// find the pivot first
	a, b := 0, len(array)
	for a < b {
		p := (a + b) / 2
		if v > array[p] {
			if v > array[b-1] {
				b = p
			} else {
				a = p + 1
			}
		} else if v < array[p] {
			if v < array[a] {
				a = p + 1
			} else {
				b = p
			}
		} else {
			return p
		}
	}
	return -1
}
