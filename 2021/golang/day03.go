package golang

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func Day03(rawInput string) {
	fmt.Print("This is day 03\n\n")

	var inputs []string

	var cols int

	for _, value := range strings.Split(rawInput, "\n") {
		//i, _ := strconv.Atoi(value)

		if cols == 0 {
			cols = len(value)
		}

		if len(value) == cols {
			inputs = append(inputs, value)
		}
	}

	var partOne, partTwo int

	diagnosticParts := [][]byte{{}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}}

	for _, row := range inputs {
		for col, bit := range strings.Split(row, "") {
			bitInt, _ := strconv.Atoi(bit)
			diagnosticParts[col] = append(diagnosticParts[col], byte(bitInt))
		}
	}

	//fmt.Println(diagnosticParts)

	var gammaRate, epsilonRate string
	for _, bits := range diagnosticParts {
		countNull := bytes.Count(bits, []byte{0})
		countOne := bytes.Count(bits, []byte{1})

		if countNull > countOne {
			gammaRate += "0"
			epsilonRate += "1"
		} else {
			gammaRate += "1"
			epsilonRate += "0"
		}
	}
	//fmt.Println(gammaRate, epsilonRate)

	gamma, _ := strconv.ParseInt(gammaRate, 2, 64)
	epsilon, _ := strconv.ParseInt(epsilonRate, 2, 64)

	partOne = int(gamma * epsilon)

	fmt.Printf("Part one solution: %v\n", partOne)
	// took 32 min
	// executed in 517.17µs (Ryzen 5 3600)
	fmt.Printf("Part two solution: %v\n", partTwo)
	// took - more minutes
	// executed in -µs (Ryzen 5 3600)
}
