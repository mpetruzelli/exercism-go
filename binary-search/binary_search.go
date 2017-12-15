package binarysearch

//iterative version
func SearchInts(array []int, target int) int {

	upper := len(array) - 1
	lower := 0
	mid := 0
	for lower <= upper {

		mid = (lower + upper) / 2

		if target == array[mid] {
			return mid
		}
		if target < array[mid] {
			upper = mid - 1
		} else {
			lower = mid + 1
		}
	}
	return -1
}
