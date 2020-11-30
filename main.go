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

	fmt.Print("Paste input: ")
	input, _ := reader.ReadString('\n')
	fmt.Println()
	/* */

	/* Development usage * /
	day := "1\n"
	input := "123 345 634"
	/* */

	switch day {
	case "1\n":
		aoc.Day01(input)
	//case "2\n":
	//	aoc.Day02()
	default:
		fmt.Println("Not found!")
	}
}
