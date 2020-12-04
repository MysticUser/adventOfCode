package main

import (
	aoc "./golang"
	"bufio"
	"fmt"
	"os"
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
		day = "4\n"
		//input = "eyr:1972 cid:100\nhcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926\n\niyr:2019\nhcl:#602927 eyr:1967 hgt:170cm\necl:grn pid:012533040 byr:1946\n\nhcl:dab227 iyr:2012\necl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277\n\nhgt:59cm ecl:zzz\neyr:2038 hcl:74454a iyr:2023\npid:3556412378 byr:2007"
		input = "pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980\nhcl:#623a2f\n\neyr:2029 ecl:blu cid:129 byr:1989\niyr:2014 pid:896056539 hcl:#a97842 hgt:165cm\n\nhcl:#888785\nhgt:164cm byr:2001 iyr:2015 cid:88\npid:545766238 ecl:hzl\neyr:2022\n\niyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719"
		/* */
	}

	switch day {
	case "1\n":
		aoc.Day01(input)
	case "2\n":
		aoc.Day02(input)
	case "3\n":
		aoc.Day03(input)
	case "4\n":
		aoc.Day04(input)
	default:
		fmt.Println("Not found!")
	}
}
