package golang

import (
	"fmt"
	"strconv"
	"strings"
)

type programStruct struct {
	opp string
	arg int
}

func Day08(rawInput string) {
	fmt.Print("This is day 08 code:\n\n")
	var matchCountPartOne, matchCountPartTwo int

	var inputRows = strings.Split(rawInput, "\n")

	var program []programStruct

	for _, row := range inputRows {
		row = strings.Trim(row, " \t")
		if row == "" {
			continue
		}

		inputs := strings.Split(row, " ")
		if len(inputs) < 2 {
			fmt.Println("Inputs fail")
			break
		}

		code := inputs[0]
		value, _ := strconv.Atoi(inputs[1])

		program = append(program, programStruct{opp: code, arg: value})
	}

	var ranInstructions []int
	var exitCode int8
	matchCountPartOne, exitCode, ranInstructions = runProgram(program)

	var fixInstruction = 1
	var newProgram = make([]programStruct, len(program))

	for exitCode == 2 {
		if fixInstruction > len(program) {
			break
		}

		copy(newProgram, program)

		// FIXING relevant instructions from last executed
		fixIndex := ranInstructions[len(ranInstructions)-fixInstruction]
		var newOpp string
		if newProgram[fixIndex].opp == "jmp" {
			newOpp = "nop"
		} else if newProgram[fixIndex].opp == "nop" {
			newOpp = "jmp"
		} else {
			fixInstruction++
			continue
		}

		//fmt.Printf("Changed line %v from %v to %v\n", fixIndex, newProgram[fixIndex].opp, newOpp)
		newProgram[fixIndex] = programStruct{opp: newOpp, arg: newProgram[fixIndex].arg}

		matchCountPartTwo, exitCode, _ = runProgram(newProgram)
		fixInstruction++
	}

	fmt.Println("Accumulator value (Part 1): ", matchCountPartOne)
	fmt.Println("Accumulator value (Part 2): ", matchCountPartTwo)
	// took 80 minutes
	// execution-time: 4.617ms (i5-8250U)
}

func runProgram(program []programStruct) (int, int8, []int) {

	var ranInstructions []int
	var exitCode int8
	var instruction, accumulator = 0, 0
	exitCode = 0

	for true {
		if containsInt(ranInstructions, instruction) {
			exitCode = 2
			break // stop infinite loop
		}
		ranInstructions = append(ranInstructions, instruction)

		if instruction >= len(program) {
			exitCode = 1
			break // successful execute
		}

		switch program[instruction].opp {
		case "acc":
			accumulator += program[instruction].arg
			instruction++
		case "jmp":
			instruction += program[instruction].arg
		case "nop":
			instruction++
		}
	}

	return accumulator, exitCode, ranInstructions
}
