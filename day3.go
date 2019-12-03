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
	X     int
	Y     int
	Steps int
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

	// Brute force intersection
	intersection := make([]Point, 0)
	for _, point1 := range points1 {
		for _, point2 := range points2 {
			if point1.X == point2.X && point1.Y == point2.Y {
				point := Point{point1.X, point1.Y, point1.Steps + point2.Steps}
				intersection = append(intersection, point)
			}
		}
	}

	shortestDistance := manhattanDistance(intersection[0])
	for _, point := range intersection[1:] {
		distance := manhattanDistance(point)
		if distance < shortestDistance {
			shortestDistance = distance
		}
	}

	fewestSteps := intersection[0].Steps
	for _, point := range intersection[1:] {
		if point.Steps < fewestSteps {
			fewestSteps = point.Steps
		}
	}

	fmt.Printf("Shortest Manhattan distance: %.0f\n", shortestDistance)
	fmt.Printf("Fewest Steps: %d\n", fewestSteps)
}

func calculatePoints(wire []string) []Point {
	points := make([]Point, 0)

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
			point := Point{x, y, steps}
			points = append(points, point)
		}
	}

	return points
}

func manhattanDistance(point Point) float64 {
	return math.Abs(float64(point.X)) + math.Abs(float64(point.Y))
}
