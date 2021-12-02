package golang

import (
	"fmt"
	"strconv"
	"strings"
)

type Movement struct {
	direction string
	move      int
}

func Day02(rawInput string) {
	fmt.Print("This is day 02\n\n")

	var inputs []Movement
	var partOne, partTwo int

	for _, value := range strings.Split(rawInput, "\n") {
		inp := strings.Split(value, " ")
		if len(inp) == 2 {
			moveCount, _ := strconv.Atoi(inp[1])

			inputs = append(inputs, Movement{inp[0], moveCount})
		}
	}

	var horizontalPos, depth int
	var horizontalPosP2, depthP2, aim int

	for _, row := range inputs {
		switch row.direction {
		case "forward":
			horizontalPos += row.move

			horizontalPosP2 += row.move
			depthP2 += aim * row.move
		case "down":
			depth += row.move

			aim += row.move
		case "up":
			depth -= row.move

			aim -= row.move
		}
	}

	partOne = horizontalPos * depth
	partTwo = horizontalPosP2 * depthP2

	fmt.Printf("Part one solution: %v\n", partOne)
	// took 18 min
	// executed in 229.402µs (Ryzen 5 3600)
	fmt.Printf("Part two solution: %v\n", partTwo)
	// took 5 more minutes
	// executed in 221.322µs (Ryzen 5 3600)
}
