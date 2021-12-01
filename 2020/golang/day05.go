package golang

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func Day05(rawInput string) {
	fmt.Print("This is day 05 code:\n\n")
	var matchCountPartOne, matchCountPartTwo int

	var inputRows = strings.Split(rawInput, "\n")
	var seatIds []int

	for _, row := range inputRows {
		if row == "" || len(row) != 10 {
			continue
		}

		code := strings.Split(row, "")

		var seatLow, seatHigh = 0.0, 127.0

		// Find row
		for i := 0; i < 7; i++ {
			var half = math.Round((seatHigh - seatLow) / 2.0)

			switch code[i] {
			case "F":
				seatHigh -= half
			case "B":
				seatLow += half
			}
		}

		// Find column
		var seatCol int
		var seatLeft, seatRight = 0.0, 7.0
		for i := 0; i < 2; i++ {
			var half = math.Round((seatRight - seatLeft) / 2.0)

			switch code[7+i] {
			case "L":
				seatRight -= half
			case "R":
				seatLeft += half
			}
		}

		if code[9] == "L" {
			seatCol = int(seatLeft)
		} else if code[9] == "R" {
			seatCol = int(seatRight)
		}

		var seatRow int
		if seatHigh == seatLow {
			seatRow = int(seatLow)
		}

		seatId := seatRow*8 + seatCol

		if matchCountPartOne < seatId {
			matchCountPartOne = seatId
		}

		seatIds = append(seatIds, seatId)
	}

	sort.Ints(seatIds)

	// Find the missing number
	if len(seatIds) > 0 {

		var lastId = seatIds[0] - 1
		for _, seatId := range seatIds {
			if lastId != seatId-1 {
				// assuming this only happens once
				matchCountPartTwo = seatId - 1
			}
			lastId = seatId
		}
	}

	fmt.Println("Valid passports part one: ", matchCountPartOne)
	fmt.Println("Valid passports part two: ", matchCountPartTwo)
	// took 52 minutes
	// executed in 317.962µs (Ryzen 5 3600)
}
