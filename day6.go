package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type object struct {
	name     string
	parents  []*object
	children []*object
}

func day6() {
	content, err := ioutil.ReadFile("input/day6.txt")
	if err != nil {
		log.Fatal(err)
	}

	inputs := strings.Split(string(content), "\n")
	objects := make(map[string]*object)

	for _, input := range inputs {
		value := strings.Split(input, ")")

		parentName := value[0]
		childName := value[1]

		var parent *object
		if val, ok := objects[parentName]; ok {
			parent = val
		} else {
			parent = &object{parentName, make([]*object, 0), make([]*object, 0)}
			objects[parent.name] = parent
		}

		var child *object
		if val, ok := objects[childName]; ok {
			child = val
		} else {
			child = &object{childName, make([]*object, 0), make([]*object, 0)}
			objects[child.name] = child
		}

		parent.children = append(parent.children, child)
		child.parents = append(child.parents, parent)
	}

	checksum := countOrbits(objects["COM"], 0)

	fmt.Printf("Orbits: %d\n", checksum)

	search(objects["YOU"], make(map[*object]bool), 0)
}

func countOrbits(object *object, previous int) int {
	if len(object.children) == 0 {
		return previous
	}

	total := previous
	for _, child := range object.children {
		total += countOrbits(child, previous+1)
	}

	return total
}

func search(object *object, visited map[*object]bool, previous int) {
	if object.name == "SAN" {
		fmt.Printf("Distance: %d\n", previous)
	}

	visited[object] = true

	for _, child := range object.children {
		if !visited[child] {
			search(child, visited, previous+1)
		}
	}

	for _, parent := range object.parents {
		if !visited[parent] {
			search(parent, visited, previous+1)
		}
	}
}
