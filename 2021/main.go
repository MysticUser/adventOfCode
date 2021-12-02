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
	//useNavigator = false

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
		day = "2\n"
		input = "forward 5\ndown 5\nforward 8\nup 3\ndown 8\nforward 2\n"
		/* */
	}

	start := time.Now()
	switch day {
	case "1\n":
		aoc.Day01(input)
	case "2\n":
		aoc.Day02(input)
	default:
		fmt.Println("Not found!")
	}
	fmt.Printf("\nExecution took %s\n", time.Since(start))
}
