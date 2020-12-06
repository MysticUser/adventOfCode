package main

import (
	aoc "./golang"
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println()

	var day, input string
	var useNavigator bool

	useNavigator = true

	if useNavigator {
		/* NAVIGATOR */
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter a day: ")
		day, _ = reader.ReadString('\n')

		fmt.Println("Paste input (press tab and enter to continue):")
		input, _ = reader.ReadString('\t')
		fmt.Println()
		/* */
	} else {
		/* Development usage */
		day = "6\n"
		input = "abc\n\na\nb\nc\n\nab\nac\n\na\na\na\na\n\nb\n"
		/* */
	}

	start := time.Now()
	switch day {
	case "1\n":
		aoc.Day01(input)
	case "2\n":
		aoc.Day02(input)
	case "3\n":
		aoc.Day03(input)
	case "4\n":
		aoc.Day04(input)
	case "5\n":
		aoc.Day05(input)
	case "6\n":
		aoc.Day06(input)
	default:
		fmt.Println("Not found!")
	}
	fmt.Printf("\nExecution took %s\n", time.Since(start))
}
