package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func day5() {
	content, err := ioutil.ReadFile("input/day5.txt")
	if err != nil {
		log.Fatal(err)
	}

	inputs := strings.Split(string(content), ",")
	memory := make([]int, len(inputs))

	for index, input := range inputs {
		value, _ := strconv.Atoi(strings.TrimSpace(input))
		memory[index] = value
	}

	runDiagnosticProgram(memory)
}

func runDiagnosticProgram(memory []int) int {
	const additionCode = 1
	const multiplicationCode = 2
	const inputCode = 3
	const outputCode = 4
	const jumpIfTrueCode = 5
	const jumpIfFalseCode = 6
	const lessThanCode = 7
	const equalsCode = 8
	const haltCode = 99

	count := 0
	instructionPointer := 0
	opcode := getOpcode(memory, instructionPointer)

	for opcode != haltCode {
		if opcode == additionCode {
			instructionPointer = addition(memory, instructionPointer)
		} else if opcode == multiplicationCode {
			instructionPointer = multiplication(memory, instructionPointer)
		} else if opcode == inputCode {
			instructionPointer = input(memory, instructionPointer)
		} else if opcode == outputCode {
			instructionPointer = output(memory, instructionPointer)
		} else if opcode == jumpIfTrueCode {
			instructionPointer = jumpIfTrue(memory, instructionPointer)
		} else if opcode == jumpIfFalseCode {
			instructionPointer = jumpIfFalse(memory, instructionPointer)
		} else if opcode == lessThanCode {
			instructionPointer = lessThan(memory, instructionPointer)
		} else if opcode == equalsCode {
			instructionPointer = equals(memory, instructionPointer)
		} else {
			fmt.Printf("Unknown opcode: %d\n", opcode)
		}

		opcode = getOpcode(memory, instructionPointer)
		count++
	}

	return memory[0]
}

func getOpcode(memory []int, instructionPointer int) int {
	instruction := memory[instructionPointer]

	return instruction % 100
}

func getInputParameters(memory []int, instructionPointer int, numberOfParameters int) []int {
	parameters := make([]int, numberOfParameters)

	instruction := memory[instructionPointer]

	mod := 1000
	div := 100

	const positionMode = 0
	const immediateMode = 1

	for i := 0; i < numberOfParameters; i++ {
		mode := (instruction % mod) / div

		if mode == positionMode {
			parameterAddress := memory[instructionPointer+i+1]
			parameters[i] = memory[parameterAddress]
		} else if mode == immediateMode {
			parameters[i] = memory[instructionPointer+i+1]
		} else {
			fmt.Printf("Unknown mode: %d\n", mode)
		}

		mod *= 10
		div *= 10
	}

	return parameters
}

func addition(memory []int, instructionPointer int) int {
	inputParameters := getInputParameters(memory, instructionPointer, 2)
	outputAddress := memory[instructionPointer+3]

	memory[outputAddress] = inputParameters[0] + inputParameters[1]

	return instructionPointer + 4
}

func multiplication(memory []int, instructionPointer int) int {
	inputParameters := getInputParameters(memory, instructionPointer, 2)
	outputAddress := memory[instructionPointer+3]

	memory[outputAddress] = inputParameters[0] * inputParameters[1]

	return instructionPointer + 4
}

func input(memory []int, instructionPointer int) int {
	outputAddress := memory[instructionPointer+1]
	const id = 5

	memory[outputAddress] = id

	return instructionPointer + 2
}

func output(memory []int, instructionPointer int) int {
	inputParameters := getInputParameters(memory, instructionPointer, 1)

	fmt.Printf("Output: %d\n", inputParameters[0])

	return instructionPointer + 2
}

func jumpIfTrue(memory []int, instructionPointer int) int {
	inputParameters := getInputParameters(memory, instructionPointer, 2)

	if inputParameters[0] != 0 {
		return inputParameters[1]
	}

	return instructionPointer + 3
}

func jumpIfFalse(memory []int, instructionPointer int) int {
	inputParameters := getInputParameters(memory, instructionPointer, 2)

	if inputParameters[0] == 0 {
		return inputParameters[1]
	}

	return instructionPointer + 3
}

func lessThan(memory []int, instructionPointer int) int {
	inputParameters := getInputParameters(memory, instructionPointer, 2)
	outputAddress := memory[instructionPointer+3]

	if inputParameters[0] < inputParameters[1] {
		memory[outputAddress] = 1
	} else {
		memory[outputAddress] = 0
	}

	return instructionPointer + 4
}

func equals(memory []int, instructionPointer int) int {
	inputParameters := getInputParameters(memory, instructionPointer, 2)
	outputAddress := memory[instructionPointer+3]

	if inputParameters[0] == inputParameters[1] {
		memory[outputAddress] = 1
	} else {
		memory[outputAddress] = 0
	}

	return instructionPointer + 4
}
