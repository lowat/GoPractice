package main

func GetNthFib(n int) int {
	prevValue := 1
	prevPrevValue := 0

	for i := 2; i < n; i++ {
		temp := prevPrevValue
		prevPrevValue = prevValue
		prevValue = prevValue + temp
	}
	if n <= 1 {
		return prevPrevValue
	}
	return prevValue
}
