package golang

import (
	"fmt"
	"strconv"
	"strings"
)

type bagRuleStruct struct {
	identifier string
	amount     int
}

func Day07(rawInput string) {
	fmt.Print("This is day 07 code:\n\n")
	var matchCountPartOne, matchCountPartTwo int

	var inputRows = strings.Split(rawInput, "\n")

	var newBagRules = make(map[string][]bagRuleStruct)
	var bagColor, bagContents string

	for _, row := range inputRows {
		row = strings.Trim(row, " \t")
		if row == "" {
			continue
		}

		rules := strings.Split(row, " ")
		bagColor = fmt.Sprint(rules[0], " ", rules[1])
		bagContents = strings.Join(rules[4:], " ")

		for _, rule := range strings.Split(bagContents, ", ") {
			ruleBagColor := strings.Split(rule, " ")
			if ruleBagColor[0] != "no" {
				numContents, err := strconv.Atoi(ruleBagColor[0])
				if err == nil {
					newBagRules[bagColor] = append(newBagRules[bagColor], bagRuleStruct{
						identifier: fmt.Sprint(ruleBagColor[1], " ", ruleBagColor[2]),
						amount:     numContents,
					})
				}
			} else {
				newBagRules[bagColor] = append(newBagRules[bagColor], bagRuleStruct{
					identifier: fmt.Sprint(ruleBagColor[1], " ", ruleBagColor[2]),
					amount:     0,
				})
			}
		}
	}

	// PART ONE
	for parentBag, bags := range newBagRules {
		if parentBag != "shiny gold" {
			if checkForShiny(bags, newBagRules) {
				matchCountPartOne++
			}
		}
	}

	// PART TWO
	matchCountPartTwo = countBags("shiny gold", newBagRules)

	fmt.Println("Matching rules (Part 1): ", matchCountPartOne)
	fmt.Println("Matching rules (Part 2): ", matchCountPartTwo)
	// part one took 95 minutes
	// part two took 61 minutes
	// took 156 minutes
	// execution-time: 9.1717ms
}

func checkForShiny(bags []bagRuleStruct, bagMap map[string][]bagRuleStruct) bool {
	for _, bag := range bags {
		if bag.identifier == "shiny gold" {
			return true
		} else {
			if bagMap[bag.identifier] != nil {
				if checkForShiny(bagMap[bag.identifier], bagMap) {
					return true
				}
			}
		}
	}

	return false
}

func countBags(checkParentBag string, bagNumbers map[string][]bagRuleStruct) int {
	var returnInt int

	for parentBag, bagsStruct := range bagNumbers {
		if parentBag == checkParentBag {
			for _, bagStruct := range bagsStruct {
				returnInt += bagStruct.amount                                 // add self-contained bags
				childContained := countBags(bagStruct.identifier, bagNumbers) // add the children bags

				returnInt += bagStruct.amount * childContained
			}
		}
	}
	return returnInt
}
