package golang

import (
	"fmt"
	"strconv"
	"strings"
)

type gameNumsStruct struct {
	number int
	amount int
}

func Day15(rawInput string) {
	fmt.Print("This is day 15 code:\n\n")
	var solutionPartOne, solutionPartTwo int

	var inputs []uint64
	for _, value := range strings.Split(rawInput, ",") {
		value = strings.Replace(value, "\t", "", -1)
		i, _ := strconv.Atoi(value)
		if i >= 0 {
			inputs = append(inputs, uint64(i))
		}
	}

	if inputs == nil {
		fmt.Println("Found no numbers")
		return
	}

	var gameNums []uint64
	var lastNum uint64

	for turn := 0; turn < 30000000; turn++ {
		var number uint64
		if len(inputs)-1 >= turn {
			number = inputs[turn]
		} else {
			number = findLastNum(gameNums)
		}

		gameNums = append(gameNums, number)
		lastNum = number

		if turn+1 == 2020 {
			solutionPartOne = int(number)
		}
	}
	solutionPartTwo = int(lastNum)

	fmt.Println("2020th number (Part 1): ", solutionPartOne)
	fmt.Println("(Part 2): ", solutionPartTwo)
	// part one took: 61 minutes - 4.681ms
	// part two took: - minutes
	// executed: -µs (Ryzen 5 3600)
}

func findLastNum(list []uint64) uint64 {
	needle := list[len(list)-1]

	for i := len(list) - 2; i >= 0; i-- {
		if list[i] == needle {
			return uint64(len(list) - (i + 1))
		}
	}

	return 0
}
