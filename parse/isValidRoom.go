package parsing

import (
	"strconv"
	"strings"
)

func IsValidRoom(str string) bool {
	arr := strings.Fields(str)
	if arr[0] == "L" || arr[0] == "l" {
		return false
	}
	_, err := strconv.Atoi(arr[1])
	if err != nil {
		return false
	}
	_, err = strconv.Atoi(arr[2])
	if err != nil {
		return false
	}
	return true;
}
