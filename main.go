package main

import (
	aoc "./golang"
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println()

	/* NAVIGATOR */
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a day: ")
	day, _ := reader.ReadString('\n')

	fmt.Println("Paste input (press tab and enter to continue):")
	input, _ := reader.ReadString('\t')
	fmt.Println()
	/* */

	/* Development usage * /
	day := "2\n"
	input := "1-3 a: abcde\n1-3 b: cdefg\n2-9 c: ccccccccc\n"
	/* */

	switch day {
	case "1\n":
		aoc.Day01(input)
	case "2\n":
		aoc.Day02(input)
	default:
		fmt.Println("Not found!")
	}
}
