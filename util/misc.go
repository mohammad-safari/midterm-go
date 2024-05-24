package util

import "strconv"

func ConvertStr2Int(input string) (result int64, err error) {
	return strconv.ParseInt(input, 10, 64)
}
