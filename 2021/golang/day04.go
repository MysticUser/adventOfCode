package golang

import (
	"fmt"
	"strconv"
	"strings"
)

type Board struct {
	board   [5][5]int
	marked  [5][5]int
	matches []int
}

func Day04(rawInput string) {
	fmt.Print("This is day 04\n\n")

	var partOne, partTwo int

	var bingoNums []string
	var boards []Board

	for entity, value := range strings.Split(rawInput, "\n\n") {

		if entity == 0 {
			bingoNums = strings.Split(value, ",")
		} else {
			var bBoard [5][5]int
			for row, numberStr := range strings.Split(value, "\n") {

				var rowInts [5]int
				var colCount int
				for _, rowString := range strings.Split(numberStr, " ") {
					number, e := strconv.Atoi(rowString)
					if e == nil {
						rowInts[colCount] = number
						colCount++
					}
				}
				if len(numberStr) > 2 {
					bBoard[row] = rowInts
				}
			}

			newBoard := Board{
				board: bBoard,
			}
			boards = append(boards, newBoard)
		}
	}

	//fmt.Println(boards)

	var bingoBoard int
	bingo := -1
	for round, bingoNumStr := range bingoNums {
		bingoNum, _ := strconv.Atoi(bingoNumStr)
		//fmt.Printf("\n%v: %v\n", round, bingoNum)
		for bKey, b := range boards {

			for row, rowNums := range b.board {
				for col, num := range rowNums {
					if bingoNum == num {
						boards[bKey].marked[row][col] = 1
						boards[bKey].matches = append(boards[bKey].matches, bingoNum)
						//fmt.Printf("Match for board %v\n", bKey)
					}
				}
			}

		}

		// Check for winners
		if round >= 5-1 {

			for bKey, b := range boards {
				var rowMatches [5]int
				var colMatches [5]int

				for row, rowNums := range b.marked {
					for col, num := range rowNums {
						if num == 1 {
							rowMatches[row]++
							colMatches[col]++
						}
						if rowMatches[row] >= 5 || colMatches[col] >= 5 {
							//fmt.Printf("\n\tBINGO for board %v, in round %v, with the number %v\n\n", bKey, round+1, bingoNum)
							bingoBoard = bKey
							bingo = bingoNum
						}
					}
				}
				//fmt.Println(bKey, rowMatches, colMatches)
			}
		}

		if bingo >= 0 {
			break
		}
	}

	// Calculate winning score
	var score int
	for row, rowNums := range boards[bingoBoard].marked {
		for col, num := range rowNums {
			if num == 0 {
				score += boards[bingoBoard].board[row][col]
			}
		}
	}

	partOne = score * bingo

	fmt.Printf("Part one solution: %v\n", partOne)
	// took 71 min
	// executed in 764.342µs (Ryzen 5 3600)
	fmt.Printf("Part two solution: %v\n", partTwo)
	// took - more minutes
	// executed in -µs (Ryzen 5 3600)
}
