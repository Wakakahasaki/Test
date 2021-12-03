package main

/**
 *
 * @param number int整型
 * @return int整型
 */
func jumpFloor(number int) int {
	// write code here
	if number <= 0 {
		return 0
	}
	if number == 1 {
		return 1
	}
	pre, cur, next := 0, 1, 1
	for i := 0; i < number; i++ {
		next = pre + cur
		pre = cur
		cur = next
	}
	return next

}
