package main

import (
	"fmt"
	"strconv"
	"strings"
)

func clearString(value string) string {
	newValue := strings.TrimSpace(value)
	newValue = strings.ReplaceAll(newValue, "\n", "")
	newValue = strings.ReplaceAll(newValue, "\r", "")
	return newValue
}

func stringToInt(value string) int64 {
	number, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0
	}
	return number
}

func stringToFloat(value string) float64 {
	number, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0.0
	}
	return number
}

func fillURL(list string) string {
	return fmt.Sprintf("https://www.imdb.com/list/%s/", list)
}
