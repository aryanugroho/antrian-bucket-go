package antrian

import "github.com/aryanugroho/antrian-bucket-go/util"

// Validate is validation util
// it's should positive number
// 1 - 5
func Validate(param string) bool {
	isValid, num := util.IsPositiveNumber(param)
	if !isValid {
		return false
	}

	if num < 1 {
		return false
	}

	if num > 5 {
		return false
	}

	return true
}
