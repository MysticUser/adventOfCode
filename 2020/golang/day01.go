package golang

import (
	"fmt"
	"strconv"
	"strings"
)

func Day01(rawInput string) {
	fmt.Print("This is day 01 code:\n\n")

	var inputs []int

	for _, value := range strings.Split(rawInput, "\n") {
		i, _ := strconv.Atoi(value)
		if i > 0 {
			inputs = append(inputs, i)
		}
	}

	var partOne, partTwo int

	for _, firstNum := range inputs {
		for _, secondNum := range inputs {
			if firstNum+secondNum == 2020 {
				//fmt.Printf("Part one solution: %v (numbers: %v, %v)\n", firstNum*secondNum, firstNum, secondNum)
				if partOne == 0 {
					partOne = firstNum * secondNum
				}
			}
			for _, thirdNum := range inputs {
				if firstNum+secondNum+thirdNum == 2020 {
					//fmt.Printf("Part two solution: %v (numbers: %v, %v and %v)\n", firstNum*secondNum*thirdNum, firstNum, secondNum, thirdNum)
					if partTwo == 0 {
						partTwo = firstNum * secondNum * thirdNum
					}
				}
			}
		}
	}

	fmt.Printf("Part one solution: %v\n", partOne)
	fmt.Printf("Part two solution: %v\n", partTwo)
	// took 33 min
	// executed in 7.966ms (Ryzen 5 3600)
}
