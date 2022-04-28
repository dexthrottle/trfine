package helper

import (
	"fmt"
	"strings"
)

// Получает кол-во знаков после запятой
func GetCountRune(number float64) int {

	numStr := fmt.Sprintf("%f", number)
	for _, v := range numStr {
		if v == '.' {
			newNumStr := strings.Split(numStr, ".")[1]
			newNumStrRune := []rune(newNumStr)
			return len(newNumStrRune)
		}
	}
	return 0
}

func StripZeroNumber(number float64) string {
	numStr := fmt.Sprintf("%f", number)
	for _, v := range numStr {
		if string(v) == ".0" {
			numStr = strings.TrimRight(strings.TrimRight(numStr, "0"), ".")
		}
	}
	return numStr
}
