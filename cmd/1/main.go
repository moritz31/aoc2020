package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

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

	rawInputData := strings.Split(string(fileBytes), "\n")

	inputData := make([]int, 0, len(rawInputData))

	for _, l := range rawInputData {
		if len(l) == 0 {
			continue
		}

		n, err := strconv.Atoi(l)
		if err != nil {
			os.Exit(1)
		}
		inputData = append(inputData, n)
	}

	result, result2 := findMatching(inputData)

	if result != -1 {
		fmt.Printf("Result: %d \n", result)
	}

	if result2 != -1 {
		fmt.Printf("Result: %d \n", result2)
		os.Exit(0)
	}

	fmt.Println("Got no matching pair")
	os.Exit(1)
}

func findMatching(input []int) (int, int) {
	firstMatching := -1
	secondMatching := -1

	for _, number := range input {
		for _, compare := range input {
			if number+compare == 2020 {
				fmt.Printf("Matching numbers found: %d - %d \n", number, compare)
				firstMatching = number * compare
			}
			for _, secondCompare := range input {
				if number+compare+secondCompare == 2020 {
					fmt.Printf("Matching triples found: %d - %d - %d \n", number, compare, secondCompare)
					secondMatching = number * compare * secondCompare
				}
			}
		}
	}
	return firstMatching, secondMatching
}
