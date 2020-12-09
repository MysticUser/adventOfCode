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
		day = "9\n"
		input = "35\n20\n15\n25\n47\n40\n62\n55\n65\n95\n102\n117\n150\n182\n127\n219\n299\n277\n309\n576\n"
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
	case "7\n":
		aoc.Day07(input)
	case "8\n":
		aoc.Day08(input)
	case "9\n":
		aoc.Day09(input)
	default:
		fmt.Println("Not found!")
	}
	fmt.Printf("\nExecution took %s\n", time.Since(start))
}
