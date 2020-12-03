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
	/* * /

	/* Development usage * /
	day := "3\n"
	input := "..##.......\n#...#...#..\n.#....#..#.\n..#.#...#.#\n.#...##..#.\n..#.##.....\n.#.#.#....#\n.#........#\n#.##...#...\n#...##....#\n.#..#...#.#"
	/* */

	switch day {
	case "1\n":
		aoc.Day01(input)
	case "2\n":
		aoc.Day02(input)
	case "3\n":
		aoc.Day03(input)
	default:
		fmt.Println("Not found!")
	}
}
