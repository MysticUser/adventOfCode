package golang

import (
	"fmt"
	"strconv"
	"strings"
)

func Day01(rawInput string) {
	fmt.Print("This is day 01\n\n")

	var inputs []int

	for _, value := range strings.Split(rawInput, "\n") {
		i, _ := strconv.Atoi(value)
		if i > 0 {
			inputs = append(inputs, i)
		}
	}

	var partOne, partTwo int

	lastNum := inputs[0]
	var newList []int
	for key, number := range inputs {
		if number > lastNum {
			partOne++
		}
		lastNum = number

		if key+2 < len(inputs) {
			newList = append(newList, inputs[key]+inputs[key+1]+inputs[key+2])
		}
	}

	lastNum = newList[0]
	for _, number := range newList {
		if number > lastNum {
			partTwo++
		}
		lastNum = number
	}

	fmt.Printf("Part one solution: %v\n", partOne)
	// took 5 min
	// executed in 65.8µs (Ryzen 5 3600)
	fmt.Printf("Part two solution: %v\n", partTwo)
	// took 17 more minutes
	// executed in 65.8µs (Ryzen 5 3600)
}
