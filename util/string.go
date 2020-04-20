package util

import "strconv"

// IsPositiveNumber will check is param on string is any positive number
func IsPositiveNumber(param string) bool {
	num, err := strconv.Atoi(param)
	if err != nil {
		return false
	}

	if num < 0 {
		return false
	}
	return true
}
