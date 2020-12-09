package golang

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Day04(rawInput string) {
	fmt.Print("This is day 04 code:\n\n")
	var matchCountPartOne, matchCountPartTwo int

	var inputRows = strings.Split(rawInput, "\n\n")

	for _, row := range inputRows {
		row = strings.Replace(row, "\n", " ", -1)

		var validKeysP1, validKeysP2 int
		var missingCID = true

		for _, passportData := range strings.Split(row, " ") {
			passportKeys := strings.Split(passportData, ":")

			if len(passportKeys) != 2 {
				continue
			}

			key := passportKeys[0]
			value := passportKeys[1]

			if key == "cid" {
				missingCID = false
			}

			if key != "" && value != "" {
				validKeysP1++

				switch key {
				case "byr":
					byr, err := strconv.Atoi(value)
					if byr < 1920 || byr > 2002 || err != nil {
						continue
					}

				case "iyr":
					iyr, err := strconv.Atoi(value)
					if iyr < 2010 || iyr > 2020 || err != nil {
						continue
					}

				case "eyr":
					eyr, err := strconv.Atoi(value)
					if eyr < 2020 || eyr > 2030 || err != nil {
						continue
					}

				case "hgt":
					if len(value) < 3 {
						continue
					}

					unit := value[len(value)-2:]
					hgt, err := strconv.Atoi(value[:len(value)-2])
					if (unit == "cm" && (hgt < 150 || hgt > 193)) || (unit == "in" && (hgt < 59 || hgt > 76) || err != nil) {
						continue
					}

				case "hcl":
					match, _ := regexp.Match("^#[0-9a-f]{6}$", []byte(value))
					if !match {
						continue
					}

				case "ecl":
					validColors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

					if !contains(validColors, value) {
						continue
					}

				case "pid":
					if len(value) != 9 {
						continue
					}

				case "cid":

				default:
					continue
				}

				validKeysP2++
			}
		}

		if validKeysP1 == 8 || (validKeysP1 == 7 && missingCID) {
			matchCountPartOne++
		}

		if validKeysP2 == 8 || (validKeysP2 == 7 && missingCID) {
			matchCountPartTwo++
		}
	}

	fmt.Println("Valid passports part one: ", matchCountPartOne)
	fmt.Println("Valid passports part two: ", matchCountPartTwo)
	// took 134 min
	// executed in 1.625ms (Ryzen 5 3600)
}
