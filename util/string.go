package util

import "strconv"

// IsPositiveNumber will check is param on string is any positive number
func IsPositiveNumber(param string) (bool, int) {
	num, err := strconv.Atoi(param)
	if err != nil {
		return false, -1
	}

	if num < 0 {
		return false, -1
	}
	return true, num
}
