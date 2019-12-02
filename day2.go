package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("input/day2.txt")
	if err != nil {
		log.Fatal(err)
	}

	inputs := strings.Split(string(content), ",")
	memory := make([]int, len(inputs))

	for index, input := range inputs {
		value, _ := strconv.Atoi(input)
		memory[index] = value
	}

	const desiredOutput = 19690720
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {

			memoryCopy := make([]int, len(memory))
			copy(memoryCopy, memory)

			output := runProgram(memoryCopy, noun, verb)

			if output == desiredOutput {
				fmt.Printf("Result: %d\n", memoryCopy[0])
				fmt.Printf("Noun: %d\n", noun)
				fmt.Printf("Verb: %d\n", verb)
				fmt.Printf("Answer: %d\n", 100*noun+verb)
				return
			}
		}
	}
}

func runProgram(memory []int, noun int, verb int) int {
	const numberOfInstructions = 4
	const additionCode = 1
	const multiplicationCode = 2
	const haltCode = 99

	// Initialize memory
	memory[1] = noun
	memory[2] = verb

	instructionPointer := 0
	opcode := memory[instructionPointer]
	for opcode != haltCode {
		input1Address := memory[instructionPointer+1]
		input2Address := memory[instructionPointer+2]

		input1 := memory[input1Address]
		input2 := memory[input2Address]

		var output int
		if opcode == additionCode {
			output = input1 + input2
		} else if opcode == multiplicationCode {
			output = input1 * input2
		} else {
			fmt.Printf("Unknown opcode: %d", opcode)
		}

		outputAddress := memory[instructionPointer+3]
		memory[outputAddress] = output

		instructionPointer += numberOfInstructions
		opcode = memory[instructionPointer]
	}

	return memory[0]
}
