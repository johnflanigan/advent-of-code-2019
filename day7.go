package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func day7() {
	content, err := ioutil.ReadFile("input/day7.txt")
	if err != nil {
		log.Fatal(err)
	}

	inputs := strings.Split(string(content), ",")
	memory := make([]int, len(inputs))

	for index, input := range inputs {
		value, err := strconv.Atoi(strings.TrimSpace(input))
		if err != nil {
			log.Fatal(err)
		}

		memory[index] = value
	}

	// Create settings
	settings := make([][]int, 0)
	for a := 0; a < 5; a++ {
		for b := 0; b < 5; b++ {
			for c := 0; c < 5; c++ {
				for d := 0; d < 5; d++ {
					for e := 0; e < 5; e++ {
						if a != b && a != c && a != d && a != e &&
							b != c && b != d && b != e &&
							c != d && c != e &&
							d != e {
							setting := []int{a, b, c, d, e}
							settings = append(settings, setting)
						}
					}
				}
			}
		}
	}

	fmt.Printf("Number of settings: %d\n", len(settings))

	maxResult := 0
	for _, setting := range settings {
		output := 0

		memoryCopy := make([]int, len(memory))
		copy(memoryCopy, memory)

		for _, element := range setting {
			memoryCopy := make([]int, len(memory))
			copy(memoryCopy, memory)

			output = runAmplificationCircuit(memoryCopy, element, output)
		}

		if output > maxResult {
			maxResult = output
		}
	}

	fmt.Printf("Max result: %d\n", maxResult)
}

func runAmplificationCircuit(memory []int, setting int, prevOutput int) int {
	const additionCode = 1
	const multiplicationCode = 2
	const inputCode = 3
	const outputCode = 4
	const jumpIfTrueCode = 5
	const jumpIfFalseCode = 6
	const lessThanCode = 7
	const equalsCode = 8
	const haltCode = 99

	instructionPointer := 0
	useSetting := true
	opcode := getOpcode(memory, instructionPointer)

	for opcode != haltCode {
		if opcode == additionCode {
			instructionPointer = addition(memory, instructionPointer)
		} else if opcode == multiplicationCode {
			instructionPointer = multiplication(memory, instructionPointer)
		} else if opcode == inputCode {

			if useSetting {
				instructionPointer = amplificationInput(memory, setting, instructionPointer)
				useSetting = false
			} else {
				instructionPointer = amplificationInput(memory, prevOutput, instructionPointer)
			}

		} else if opcode == outputCode {
			return amplificationOutput(memory, instructionPointer)
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
	}

	return memory[0]
}

func amplificationInput(memory []int, setting int, instructionPointer int) int {
	outputAddress := memory[instructionPointer+1]

	memory[outputAddress] = setting

	return instructionPointer + 2
}

func amplificationOutput(memory []int, instructionPointer int) int {
	inputParameters := getInputParameters(memory, instructionPointer, 1)
	return inputParameters[0]
}
