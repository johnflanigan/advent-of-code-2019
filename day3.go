package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

func main() {
	content, err := ioutil.ReadFile("input/day3.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(string(content), "\n")
	wire1 := strings.Split(input[0], ",")
	wire2 := strings.Split(input[1], ",")

	points1 := calculatePoints(wire1)
	points2 := calculatePoints(wire2)

	intersection := make(map[Point]int)
	for point1, steps1 := range points1 {
		if steps2, ok := points2[point1]; ok {
			intersection[point1] = steps1 + steps2
		}
	}

	var shortestDistance float64
	// Initialize shortest distance with any point from map
	for point := range intersection {
		shortestDistance = manhattanDistance(point)
		break
	}
	for point := range intersection {
		distance := manhattanDistance(point)
		if distance < shortestDistance {
			shortestDistance = distance
		}
	}

	var fewestSteps int
	// Initialize fewestSteps with any step count from map
	for _, steps := range intersection {
		fewestSteps = steps
	}
	for _, steps := range intersection {
		if steps < fewestSteps {
			fewestSteps = steps
		}
	}

	fmt.Printf("Shortest Manhattan distance: %.0f\n", shortestDistance)
	fmt.Printf("Fewest Steps: %d\n", fewestSteps)
}

func calculatePoints(wire []string) map[Point]int {
	points := make(map[Point]int)

	x := 0
	y := 0
	steps := 0

	for _, direction := range wire {
		distance, _ := strconv.Atoi(direction[1:])

		for i := 0; i < distance; i++ {
			if direction[0] == 'U' {
				y++
			} else if direction[0] == 'R' {
				x++
			} else if direction[0] == 'D' {
				y--
			} else if direction[0] == 'L' {
				x--
			}

			steps++
			point := Point{x, y}

			if _, ok := points[point]; !ok {
				points[point] = steps
			}
		}
	}

	return points
}

func manhattanDistance(point Point) float64 {
	return math.Abs(float64(point.X)) + math.Abs(float64(point.Y))
}
