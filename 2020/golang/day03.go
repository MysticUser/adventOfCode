package golang

import (
	"fmt"
	"strings"
)

func Day03(rawInput string) {
	fmt.Print("This is day 03 code:\n\n")
	var matchCountPartOne, matchCountPartTwo = 0, 1
	var slopeEncounters, posX = 0, 0

	var slopes = [5][2]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	var inputMap = strings.Split(rawInput, "\n")

	for _, slope := range slopes {
		var slopeRight, slopeDown = slope[0], slope[1]
		slopeEncounters, posX = 0, 0

		for i := 0; i < len(inputMap); i += slopeDown {
			var tempPosX = posX % len(inputMap[i])

			if string(inputMap[i][tempPosX]) == "#" {
				slopeEncounters++

				if slopeRight == 3 && slopeDown == 1 {
					matchCountPartOne++
				}
			}
			posX += slopeRight
		}

		matchCountPartTwo *= slopeEncounters
	}

	fmt.Println("Encounters part one: ", matchCountPartOne)
	fmt.Println("Encounters part two: ", matchCountPartTwo)
	// took 46 minutes
	// executed in 52.98µs (Ryzen 5 3600)
}
