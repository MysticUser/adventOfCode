package golang

import (
	"fmt"
	"strconv"
	"strings"
)

func Day09(rawInput string) {
	fmt.Print("This is day 09 code:\n\n")

	var inputs []int

	for _, value := range strings.Split(rawInput, "\n") {
		i, _ := strconv.Atoi(value)
		if i > 0 {
			inputs = append(inputs, i)
		}
	}

	var partOne, partTwo int

	var preambleLength = 25
	var preamble []int

	for i := preambleLength; i < len(inputs); i++ {
		preamble = inputs[i-preambleLength : i]
		if !verifyNum(inputs[i], preamble) {
			partOne = inputs[i]
			break
		}
	}

	var startI, testI int
	var groupI []int
	for partTwo == 0 && startI < len(inputs) {
		testI = 0
		groupI = []int{}
		for i := startI; i < len(inputs); i++ {
			testI += inputs[i]
			groupI = append(groupI, inputs[i])

			if testI == partOne {
				partTwo = highest(groupI) + lowest(groupI)
				break
			} else if testI > partOne {
				startI++
				break
			}
		}
	}

	fmt.Printf("Part one solution: %v\n", partOne)
	fmt.Printf("Part two solution: %v\n", partTwo)
	// took 42 min
	// execution 1.303ms (Ryzen 5 3600)
}

func verifyNum(num int, preamble []int) bool {

	var matches []int

	for _, firstNum := range preamble {
		for _, secondNum := range preamble {
			if firstNum != secondNum && firstNum+secondNum == num {
				matches = append(matches, num)
			}
		}
	}

	if len(matches) > 0 {
		return true
	}

	return false
}
