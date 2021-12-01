package golang

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Day10(rawInput string) {
	fmt.Print("This is day 10 code:\n\n")

	var inputs []int

	for _, value := range strings.Split(rawInput, "\n") {
		i, _ := strconv.Atoi(value)
		if i > 0 {
			inputs = append(inputs, i)
		}
	}

	if len(inputs) == 0 {
		fmt.Println("No input")
		return
	}

	var partOne, partTwo int

	sort.Ints(inputs)

	// PART ONE
	var lastNum, oneJumps, threeJumps int
	for _, num := range inputs {
		difference := num - lastNum
		if difference == 3 {
			threeJumps++
		} else if difference == 1 {
			oneJumps++
		}
		//fmt.Println(num, num-lastNum)
		lastNum = num
	}
	threeJumps++
	partOne = oneJumps * threeJumps

	// PART TWO
	maxJoltage := inputs[len(inputs)-1] + 3
	inputs = append(inputs, maxJoltage)

	var solutions = make(map[int][]int)
	solutions[0] = []int{inputs[0]}
	//var solutions = [][]int{{0, inputs[0]}}

	for _, num := range inputs {
		//if key == 40 { break }
		//fmt.Print("Num: ",num)

		var matches []int
		for i := 1; i <= 3; i++ {
			if containsInt(inputs, num+i) {
				matches = append(matches, i)
			}
		}
		//fmt.Println(" - matches:", matches)

		//fmt.Println("SolutionsLen", len(solutions))
		var newCopiedKeys = [][]int{{}, {}}
		if len(matches) > 1 {
			//fmt.Println("Adding more solutions..")
			var copies []int
			for solKey, sol := range solutions {
				if sol[len(sol)-1] == num {
					copies = append(copies, solKey)
				}
			}
			for _, solKey := range copies {

				for j := 0; j < len(matches)-1; j++ {
					var newSolution []int
					copy(newSolution, solutions[solKey])
					newKey := len(solutions)
					solutions[newKey] = solutions[solKey]
					newCopiedKeys[j] = append(newCopiedKeys[j], newKey)
				}

			}
		}
		//fmt.Println("SolutionsLen after", len(solutions))

		for solKey, sol := range solutions {
			//fmt.Print(solKey, sol)

			joltage := 0
			if len(sol) > 0 {
				joltage = sol[len(sol)-1]
			}

			//fmt.Print(" ",joltage)

			if num == joltage && len(matches) > 0 {
				var nextNum int
				if containsInt(newCopiedKeys[1], solKey) {
					nextNum = joltage + 3
				} else if containsInt(newCopiedKeys[0], solKey) {
					nextNum = joltage + 2
				} else if len(matches) > 1 {
					nextNum = joltage + 1
				} else {
					nextNum = joltage + matches[0]
				}
				solutions[solKey] = append(solutions[solKey], nextNum)
				//fmt.Println(" - added: ", nextNum)

			}
		}

		//fmt.Println()
		//fmt.Println()
	}
	partTwo = len(solutions)

	//for _, solutio := range solutions {
	//	fmt.Println(solutio)
	//}

	fmt.Printf("Part one solution: %v\n", partOne)
	fmt.Printf("Part two solution: %v\n", partTwo)
	// part one took 9 min, 123.9µs
	// part two took almost 4 hours..
	// took - min
	// execution - (Ryzen 5 3600)
}
