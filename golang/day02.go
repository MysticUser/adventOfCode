package golang

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Day02(rawInput string) {
	fmt.Print("This is day 02 code:\n\n")

	var matchCountPartOne, matchCountPartTwo int

	for _, value := range strings.Split(rawInput, "\n") {
		// Check if line matches pattern to check that all values are there
		regMatch, _ := regexp.Match("([0-9]+)-([0-9]+) [a-z]: [a-z]+", []byte(value))
		if !regMatch && value != "" {
			break
		}

		inputRow := strings.Split(value, ": ")
		pw := inputRow[1]

		inputSettings := strings.Split(inputRow[0], " ")
		pwRange := inputSettings[0]
		pwLetter := inputSettings[1]

		pwRangeParts := strings.Split(pwRange, "-")
		pwRangeLow, _ := strconv.Atoi(pwRangeParts[0])
		pwRangeHigh, _ := strconv.Atoi(pwRangeParts[1])

		pwCount := strings.Count(pw, pwLetter)

		if pwCount >= pwRangeLow && pwCount <= pwRangeHigh {
			matchCountPartOne++
		}

		firstMatch := string(pw[pwRangeLow-1]) == pwLetter
		secondMatch := string(pw[pwRangeHigh-1]) == pwLetter

		if (firstMatch && !secondMatch) || (!firstMatch && secondMatch) {
			matchCountPartTwo++
		}
	}

	fmt.Println("Matches part one: ", matchCountPartOne)
	fmt.Println("Matches part two: ", matchCountPartTwo)
	// clock failed
}
