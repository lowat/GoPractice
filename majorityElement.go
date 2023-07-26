package main

func MajorityElement(array []int) int {
	ans := 0
	count := 0
	for _, num := range array {
		if count == 0 {
			ans = num
		}
		if ans == num {
			count += 1
		} else {
			count -= 1
		}
	}
	return ans
}
