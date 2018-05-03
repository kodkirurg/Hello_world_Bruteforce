package main

import (
	"fmt"
	"runtime"
	"os"
	"os/exec"
	//"strconv"
	//"time"
)
func main(){
	//var s string = runtime.GOOS
	clearTerminal()
	fmt.Printf("%d\n",int(generateCharArray()[0]))
}

func bruteForce(){

}

//const untyped, var typed

//returns -1 if not found and index if found
func binarySearch(sortedArray []int) int{
	//var index int

	return -1
}

func generateCharArray() []int {
	const lowerCharValue,upperCharValue int = 32, 122
	var sortedCharArray []int
	sortedCharArray = make([]int, upperCharValue-lowerCharValue+1)
	for x:= 0 ; x <= upperCharValue-lowerCharValue ; x++ {
		sortedCharArray[x] = x+lowerCharValue
	}
	return sortedCharArray
}

func clearTerminal(){
	var consoleCommand string
	switch system:=runtime.GOOS; system{
	case "linux" : 	
		consoleCommand = "clear"
	case "windows" : 
		consoleCommand = "cls"
	default:
		panic("Unsupported OS")
	}
	c := exec.Command(consoleCommand)
	c.Stdout = os.Stdout
	c.Run()
}
