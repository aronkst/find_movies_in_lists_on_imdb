package main

import (
	"fmt"
	"os"
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

func fillURLWithList(list string) string {
	if strings.Contains(list, "//") {
		return list
	}
	return fmt.Sprintf("https://www.imdb.com/list/%s/", list)
}

func getAverage(score float64, metascore int64) float64 {
	if metascore == 0 {
		return score
	}
	metascoreFloat := float64(metascore) / 10.0
	return (score + metascoreFloat) / 2.0
}

func getParamsValues() (string, float64, int64) {
	list := os.Args[1]
	average := stringToFloat(os.Args[2])
	votes := stringToInt(os.Args[3])
	return list, average, votes
}
