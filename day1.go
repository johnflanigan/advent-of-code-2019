package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("input/day1.txt")
	if err != nil {
		log.Fatal(err)
	}

	inputs := strings.Split(string(content), "\n")

	totalFuel := 0

	for _, input := range inputs {
		mass, _ := strconv.Atoi(input)

		totalFuel += calculateFuelRequired(mass)
	}

	fmt.Printf("Fuel required: %d", totalFuel)
}

func calculateFuelRequired(mass int) int {
	fuel := (mass / 3) - 2

	if fuel <= 0 {
		return 0
	} else {
		return fuel + calculateFuelRequired(fuel)
	}
}
