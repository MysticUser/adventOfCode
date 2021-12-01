package golang

import (
	"fmt"
	"strings"
)

func Day11(rawInput string) {
	fmt.Print("This is day 11 code:\n\n")
	var solutionPartOne, solutionPartTwo int

	var inputRows = strings.Split(rawInput, "\n")
	var seatMap = make(map[int][]string)

	for rowId, row := range inputRows {
		if row == "" || row == "\t" {
			continue
		}

		for _, char := range row {
			seatMap[rowId] = append(seatMap[rowId], string(char))
		}
	}

	for true {
		var result bool
		seatMap, result = populateSeating(seatMap)
		if result == false {
			break
		}
	}

	for _, row := range seatMap {
		for _, char := range row {
			if char == "#" {
				solutionPartOne++
			}
		}
	}

	fmt.Println("Available seats (part one): ", solutionPartOne)
	fmt.Println("Valid passports part two: ", solutionPartTwo)
	// part one took 91 minutes, 801.52ms (i5-8250U)
	// took - minutes
	// executed in -µs ()
}

func populateSeating(seats map[int][]string) (map[int][]string, bool) {
	var newSeats = make(map[int][]string)
	var changed = false

	for i := 0; i < len(seats); i++ {
		//fmt.Println(i, "\t", seats[i])

		for coll := 0; coll < len(seats[i]); coll++ {
			char := seats[i][coll]

			adjacentOcc := adjacentOccupied(seats, i, coll)
			if char == "L" && adjacentOcc == 0 {
				newSeats[i] = append(newSeats[i], "#")
				changed = true
			} else if char == "#" && adjacentOcc >= 4 {
				newSeats[i] = append(newSeats[i], "L")
				changed = true
			} else {
				newSeats[i] = append(newSeats[i], seats[i][coll])
			}
		}
	}

	return newSeats, changed
}

func adjacentOccupied(seats map[int][]string, row int, col int) int {
	var occupied int
	for rowI := 0; rowI < 3; rowI++ {
		for colI := 0; colI < 3; colI++ {
			var checkRowI, checkColI int
			switch rowI {
			case 0:
				checkRowI = row - 1
			case 1:
				checkRowI = row
			case 2:
				checkRowI = row + 1
			}

			switch colI {
			case 0:
				checkColI = col - 1
			case 1:
				checkColI = col
			case 2:
				checkColI = col + 1
			}

			if checkRowI >= 0 && checkRowI < len(seats) &&
				checkColI >= 0 && checkColI < len(seats[0]) &&
				!(checkRowI == row && checkColI == col) {
				if seats[checkRowI][checkColI] == "#" {
					occupied++
				}
			}
		}
	}

	return occupied
}
