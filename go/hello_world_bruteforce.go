package main

import (
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"sync"
	"time"
)

const helloWorld = "Hello World!"
const timeBetweenFrameUpdate = 5 //Milliseconds
const workers int = 4

var logger = log.New(os.Stdout, "", 0)
var str string
var mutex = &sync.Mutex{}

func main() {
	clearTerminal()
	var wg sync.WaitGroup
	for x := 0; x < workers; x++ {
		wg.Add(1)
		go bruteForce(&wg, x+1)

	}
	wg.Wait() //wait for workers to join
	time.Sleep(time.Second)
	logger.Printf("\033[" + strconv.Itoa(workers+1) + ";0H")
	logger.Printf("\n \n \n")
}

func bruteForce(wg *sync.WaitGroup, worker int) {
	defer wg.Done() //when finished tell main thread

	generatedSlice := generateCharSlice()
	for _, v := range helloWorld {
		index := binarySearch(generatedSlice, int(v), worker)
		if index == -1 {
			panic("Error; char missing")
		}
		str = str + string(v)
	}
}

//returns -1 if not found and index if found
func binarySearch(sortedSlice []int, element int, workerLine int) int {
	indexLower := 0
	inBetweenIndex := len(sortedSlice) / 2
	indexUpper := len(sortedSlice) - 1

	for indexLower <= indexUpper {
		printWorker(inBetweenIndex, workerLine)
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

func printWorker(charIndex int, workerLine int) {
	time.Sleep(time.Millisecond * timeBetweenFrameUpdate)
	mutex.Lock()
	var s string
	for i := range generateCharSlice() {
		if i == charIndex {
			s = s + string(generateCharSlice()[charIndex])
		} else {
			s = s + "_"
		}
	}
	logger.Printf("\033[" + strconv.Itoa(workerLine) + ";0H") //move to correct line for printing worker progress
	logger.Printf(s + "\n")
	mutex.Unlock()
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
