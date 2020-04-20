package antrian

import "github.com/aryanugroho/antrian-bucket-go/util"

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
