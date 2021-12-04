package golang

import (
	"fmt"
	"strconv"
	"strings"
)

type Board struct {
	board      [5][5]int
	marked     [5][5]int
	bingo      bool
	bingoScore int
}

func Day04(rawInput string) {
	fmt.Print("This is day 04\n\n")

	var partOne, partTwo int

	var bingoNums []string
	var boards []Board

	// Translate input to structs
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

	// Play bingo

	firstBingoBoard := -1
	var lastBingoBoard int

	for round, bingoNumStr := range bingoNums {
		bingoNum, _ := strconv.Atoi(bingoNumStr)
		for bKey, b := range boards {

			for row, rowNums := range b.board {
				for col, num := range rowNums {
					if bingoNum == num {
						boards[bKey].marked[row][col] = 1
					}
				}
			}

		}

		// Check for winners
		if round >= 5-1 {

			for bKey, b := range boards {
				if b.bingo == false {
					var rowMatches [5]int
					var colMatches [5]int

					for row, rowNums := range b.marked {
						for col, num := range rowNums {
							if num == 1 {
								rowMatches[row]++
								colMatches[col]++
							}

							// Check if BINGO
							if rowMatches[row] >= 5 || colMatches[col] >= 5 {
								//fmt.Printf("\n\tBINGO for board %v, in round %v, with the number %v\n", bKey, round+1, bingoNum)
								boards[bKey].bingo = true

								// CALCULATE SCORE
								boards[bKey].bingoScore = calculateBoardSum(b) * bingoNum

								if firstBingoBoard == -1 {
									// First win (part one)
									firstBingoBoard = bKey
								}

								// Last win (part one)
								lastBingoBoard = bKey
							}
						}
					}

					// Don't check for more bingo's if this board got the bingo
					if b.bingo == true {
						break
					}
				}
			}
		}
	}

	partOne = boards[firstBingoBoard].bingoScore
	partTwo = boards[lastBingoBoard].bingoScore

	fmt.Printf("Part one solution: %v\n", partOne)
	// took 71 min
	// executed in 764.342µs (Ryzen 5 3600)
	fmt.Printf("Part two solution: %v\n", partTwo)
	// took 59 more minutes
	// executed in 1.83535ms (Ryzen 5 3600)
}

func calculateBoardSum(board Board) int {
	//var boardString string
	var score int

	for row, rowNums := range board.marked {
		for col, num := range rowNums {
			if num == 0 {
				//boardString += fmt.Sprintf("(%02d)", board.board[row][col])
				score += board.board[row][col]
				//} else {
				//boardString += "  * "
			}
		}
		//boardString += "\n"
	}

	//fmt.Println(boardString)
	return score
}
