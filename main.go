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
		day = "10\n"
		input = "16\n10\n15\n5\n1\n11\n7\n19\n6\n12\n4\n"
		//input = "28\n33\n18\n42\n31\n14\n46\n20\n48\n47\n24\n23\n49\n45\n19\n38\n39\n11\n1\n32\n25\n35\n8\n17\n7\n9\n4\n2\n34\n10\n3\n"

		day = "12\n"
		input = "F10\nN3\nF7\nR90\nF11\n"
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
	case "10\n":
		aoc.Day10(input)
	case "11\n":
		fmt.Println("Skipped")
	case "12\n":
		aoc.Day12(input)
	default:
		fmt.Println("Not found!")
	}
	fmt.Printf("\nExecution took %s\n", time.Since(start))
}
