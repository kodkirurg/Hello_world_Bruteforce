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

const message = "Hello World! How is it going? This is a long message"
const timeBetweenFrameUpdate = 150 //Milliseconds
const workers int = 4

var logger = log.New(os.Stdout, "", 0) // thread safe print
var mutex = &sync.Mutex{}
var bruteforceString = string(make([]byte, len(message)))
var wg sync.WaitGroup

func main() {
	clearTerminal()
	helloWorldindex := 0
	var channelArray [workers]chan int
	for x := 0; x < workers; x++ { //spawn workers and assign first workload
		channelArray[x] = make(chan int, 2)
		go bruteForce(&wg, x+1, channelArray[x])
		channelArray[x] <- helloWorldindex
		helloWorldindex++
	}
	//assign equal work to workers via channels
	for ; helloWorldindex < len(message); helloWorldindex += 4 {
		for i := range channelArray {
			if !(helloWorldindex+i < len(message)) {
				break
			}
			channelArray[i] <- helloWorldindex + i
		}
	}
	//shutdown workers
	for x := 0; x < workers; x++ {
		channelArray[x] <- -1
	}
	wg.Wait() //wait for workers to join
}

func bruteForce(wg *sync.WaitGroup, worker int, channel chan int) {
	wg.Add(1)
	defer wg.Done() //when finished, notify
	generatedSlice := generateCharSlice()
	for {
		helloWorldIndex := <-channel
		if helloWorldIndex == -1 { //shutdown
			return
		}
		index := binarySearch(generatedSlice, int(message[helloWorldIndex]), worker, helloWorldIndex)
		if index == -1 {
			panic("Error; char missing")
		}
	}
}

//returns -1 if not found and index if found
func binarySearch(sortedSlice []int, element int, workerLine int, helloWorldIndex int) int {
	indexLower := 0
	inBetweenIndex := len(sortedSlice) / 2
	indexUpper := len(sortedSlice) - 1

	for indexLower <= indexUpper {
		printWorker(inBetweenIndex, workerLine, helloWorldIndex)
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

func printWorker(charIndex int, workerLine int, helloWorldIndex int) {
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
	logger.Printf("\033[" + strconv.Itoa(workerLine+1) + ";0H") //move to correct line for printing worker progress
	logger.Printf(s + "\n")
	logger.Printf("\033[0;0f") //move to correct line for printing worker progress //move to hello world string and show latest bruteforce attempt char
	bruteforceString = bruteforceString[:helloWorldIndex] + string(generateCharSlice()[charIndex]) + bruteforceString[helloWorldIndex+1:]
	logger.Printf(bruteforceString)
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
