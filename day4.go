package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func day4() {
	content, err := ioutil.ReadFile("input/day4.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(string(content), "-")
	start, _ := strconv.Atoi(strings.TrimSpace(input[0]))
	end, _ := strconv.Atoi(strings.TrimSpace(input[1]))

	satisfied := 0
	for i := start; i <= end; i++ {
		if evaluateRules(i) {
			satisfied++
		}
	}

	fmt.Printf("Number of passwords satisfying rules: %d\n", satisfied)
}

func evaluateRules(input int) bool {
	if input < 100000 || input > 999999 {
		return false
	}

	first := (input % 1000000) / 100000
	second := (input % 100000) / 10000
	third := (input % 10000) / 1000
	fourth := (input % 1000) / 100
	fifth := (input % 100) / 10
	sixth := input % 10

	if !(first == second && second != third ||
		first != second && second == third && third != fourth ||
		second != third && third == fourth && fourth != fifth ||
		third != fourth && fourth == fifth && fifth != sixth ||
		fourth != fifth && fifth == sixth) {
		return false
	}

	if !(first <= second &&
		second <= third &&
		third <= fourth &&
		fourth <= fifth &&
		fifth <= sixth) {
		return false
	}

	return true
}
