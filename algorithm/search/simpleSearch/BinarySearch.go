package simpleSearch

// BinarySearch 二分查找
func BinarySearch(arr []int, key int) int {
	low, mid, high := 0, 0, len(arr)-1

	for low <= high {
		mid = low + (high-low)/2
		if arr[mid] < key {
			low = mid + 1
		} else if arr[mid] > key {
			high = mid - 1
		} else {
			return mid
		}
	}
	return -1
}
