package golang

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Day12(rawInput string) {
	fmt.Print("This is day 12 code:\n\n")
	var solutionPartOne, solutionPartTwo int

	var inputRows = strings.Split(rawInput, "\n")

	var posY, posX int
	var facing = 90

	var wayPosY, wayPosX, wayDir = 1, 10, 90
	var partTwoPosY, partTwoPosX int

	for _, row := range inputRows {
		row = strings.Trim(row, " \t")
		if len(row) > 1 {

			action := string(row[0])
			value, _ := strconv.Atoi(row[1:])

			switch action {
			case "N":
				posY += value
				wayPosY += value
			case "S":
				posY -= value
				wayPosY -= value
			case "E":
				posX += value
				wayPosX += value
			case "W":
				posX -= value
				wayPosX -= value
			case "L":
				facing -= value

				switch value {
				case 90:
					wayPosX, wayPosY = 0-wayPosY, wayPosX
				case 180:
					wayPosX, wayPosY = 0-wayPosX, 0-wayPosY
				case 270:
					wayPosX, wayPosY = wayPosY, 0-wayPosX
				}
			case "R":
				facing += value

				switch value {
				case 90:
					wayPosX, wayPosY = wayPosY, 0-wayPosX
				case 180:
					wayPosX, wayPosY = 0-wayPosX, 0-wayPosY
				case 270:
					wayPosX, wayPosY = 0-wayPosY, wayPosX
				}

				wayDir += value
			case "F":
				switch facing {
				case 0:
					posY += value
				case 90:
					posX += value
				case 180:
					posY -= value
				case 270:
					posX -= value
				default:
					fmt.Println("Broke in facing on row: ", row, "facing: ", facing)
					return
				}

				partTwoPosX += wayPosX * value
				partTwoPosY += wayPosY * value
			default:
				fmt.Println("Broke on row: ", row)
				return
			}

			for facing < 0 || facing > 270 {
				if facing < 0 {
					facing = 360 + facing
				}
				if facing > 270 {
					facing = facing - 360
				}
			}
			//fmt.Printf("%v\tWE %v (%v)\tSN %v (%v)\n", row, partTwoPosX, wayPosX, partTwoPosY, wayPosY)
		}
	}

	solutionPartOne = mathAbs(posX) + mathAbs(posY)
	solutionPartTwo = mathAbs(partTwoPosX) + mathAbs(partTwoPosY)

	fmt.Println("Manhattan Distance (Part 1): ", solutionPartOne)
	fmt.Println("Manhattan Distance (Part 2): ", solutionPartTwo)
	// part one took: 30? minutes
	// part two took: 31 minutes
	// executed: 139.7µs (Ryzen 5 3600)
}

func mathAbs(x int) int {
	return int(math.Abs(float64(x)))
}
