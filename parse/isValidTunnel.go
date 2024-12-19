package parsing

import (
	"strconv"
	"strings"
)

func IsValidTunnel(tunnel string) bool {
	if !strings.Contains(tunnel, "-") {
		return false
	}
	arr := strings.Split(tunnel, "-")
	_, err := strconv.Atoi(arr[0])
	if err != nil {
		return false
	}
	_, err = strconv.Atoi(arr[1])
	if err != nil {
		return false
	}
	return true
}
