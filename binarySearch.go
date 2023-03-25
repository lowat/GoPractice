func BinarySearch(array []int, target int) int {
	left := 0
	right := len(array) - 1
	for left <= right {
		mid := (left + right) / 2
		midValue := array[mid]
		if target == midValue {
			return mid
		} else if target > midValue {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}