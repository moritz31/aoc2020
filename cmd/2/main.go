package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type rule struct {
	min       int
	max       int
	character string
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Printf("USAGE : %s <target_filename> \n", os.Args[0])
		os.Exit(0)
	}

	fileName := os.Args[1]

	fileBytes, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	valid := parsePasswordList(string(fileBytes))

	fmt.Printf("Found %d vaild passwords \n", valid)
}

func parsePasswordList(inputList string) int {
	validCount := 0
	lines := strings.Split(inputList, "\n")

	for _, line := range lines {
		parts := strings.Split(line, ":")
		rule := generateRule(parts[0])
		if validatePasswordWithRule(rule, parts[1]) {
			validCount = validCount + 1
		}
	}
	return validCount
}

func generateRule(ruleAsString string) rule {
	ruleParts := strings.Split(ruleAsString, " ")

	createdRule := rule{character: ruleParts[1]}

	numbers := strings.Split(ruleParts[0], "-")

	createdRule.min, _ = strconv.Atoi(numbers[0])
	createdRule.max, _ = strconv.Atoi(numbers[1])

	return createdRule
}

func validatePasswordWithRule(spec rule, password string) bool {
	// count := strings.Count(password, spec.character)
	// if count >= spec.min && count <= spec.max {
	// 	return true
	// }
	// return false
	firstMatch := false
	secondMatch := false
	fmt.Printf("%d %d %s\n", spec.min, spec.max, spec.character)
	fmt.Printf("%s\n", string(password[spec.min]))
	if string(password[spec.min]) == spec.character {
		firstMatch = true
	}

	if string(password[spec.max]) == spec.character {
		secondMatch = true
	}
	fmt.Printf("%t %t\n", firstMatch, secondMatch)
	return firstMatch != secondMatch
}
