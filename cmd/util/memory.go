package util

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseMemorySize(memory string) (int64, error) {
	var memorySize int64 = -1
	// extract memory size (number) and unit (g/G or  m/M)
	strSize := len(memory)
	strNum := memory[:strSize-1]
	num, err := strconv.Atoi(strNum)
	if err != nil {
		return memorySize, err
	}
	unit := strings.ToLower(memory[strSize-1:])
	switch unit {
	case "g":
		//GB
		memorySize = 1073741824 * int64(num)
	case "m":
		//MB
		memorySize = 1048576 * int64(num)
	default:
		err := fmt.Errorf("Unknown memory unit (%s)", unit)
		return memorySize, err
	}
	return memorySize, nil
}
