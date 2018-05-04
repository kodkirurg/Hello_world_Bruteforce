package main

import (
	"os"
	"os/exec"
	"runtime"
	//"strconv"
	"time"
)

const helloWorld = "Hello World!"

var str string = ""

func main() {
	clearTerminal()
	bruteForce()
}

func bruteForce() {
	generatedSlice := generateCharSlice()
	for _, v := range helloWorld {
		index := binarySearch(generatedSlice, int(v))
		if index == -1 {
			panic("Error; char missing")
		}
		str = str + string(v)
	}
}

func printStatus(index int) {
	time.Sleep(time.Second)
	clearTerminal()
	println(str + string(generateCharSlice()[index]))
	for i := range generateCharSlice() {
		if i == index {
			print(string(generateCharSlice()[index]))
		} else {
			print("_")
		}
	}
}

//returns -1 if not found and index if found
func binarySearch(sortedSlice []int, element int) int {
	indexLower := 0
	inBetweenIndex := len(sortedSlice) / 2
	indexUpper := len(sortedSlice) - 1

	for indexLower <= indexUpper {
		printStatus(inBetweenIndex)
		if sortedSlice[inBetweenIndex] > element { //go smaller
			indexUpper = inBetweenIndex - 1
		} else if sortedSlice[inBetweenIndex] < element { //go larger
			indexLower = inBetweenIndex + 1
		} else { //lagom
			return inBetweenIndex
		}
		inBetweenIndex = indexLower + (indexUpper-indexLower)/2
	}
	return -1
}

func generateCharSlice() []int {
	const lowerCharValue, upperCharValue int = 32, 122
	var sortedCharSlice []int
	sortedCharSlice = make([]int, upperCharValue-lowerCharValue+1)
	for x := 0; x <= upperCharValue-lowerCharValue; x++ {
		sortedCharSlice[x] = x + lowerCharValue
	}
	return sortedCharSlice
}

func clearTerminal() {
	var consoleCommand string
	switch system := runtime.GOOS; system {
	case "linux":
		consoleCommand = "clear"
	case "windows":
		consoleCommand = "cls"
	default:
		panic("Unsupported OS")
	}
	c := exec.Command(consoleCommand)
	c.Stdout = os.Stdout
	c.Run()
}
