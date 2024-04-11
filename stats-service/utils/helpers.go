package utils

import (
	"strconv"
	"strings"
)

func ParseMemInfoLine(raw string) (key string, value int) {
	text := strings.ReplaceAll(raw[:len(raw)-2], " ", "")
	keyValue := strings.Split(text, ":")
	return keyValue[0], ToInt(keyValue[1])
}

func ToInt(s string) int {
	if s == "" {
		return 0
	}

	num, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return num
}
