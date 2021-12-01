package golang

import (
	"fmt"
	"strings"
)

func Day06(rawInput string) {
	fmt.Print("This is day 06 code:\n\n")
	var matchCountPartOne, matchCountPartTwo int

	var inputRows = strings.Split(rawInput, "\n\n")

	for _, row := range inputRows {
		rowMatches := 0
		row = strings.Trim(row, " \t")
		newRow := strings.Replace(row, "\n", "", -1)

		var partySize int
		for _, value := range strings.Split(row, "\n") {
			if value != "" {
				partySize++
			}
		}

		var chars []string

		for _, char := range strings.Split(newRow, "") {
			if char != "" {
				rowMatches++
				if !contains(chars, char) {
					matchCountPartOne++
					chars = append(chars, char)

					if strings.Count(newRow, char) == partySize {
						matchCountPartTwo++
					}
				}
			}
		}
	}

	fmt.Println("Yes answers  (Part 1): ", matchCountPartOne)
	fmt.Println("Same answers (Part 2): ", matchCountPartTwo)
	// took 60+ minutes
	// executed in 1.332ms (Ryzen 5 3600)
}
