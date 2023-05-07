package util

import "strconv"

func ConvertToInt(str string) (int, error) {
	return strconv.Atoi(str)
}
