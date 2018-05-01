package main

import (
	"fmt"
	"runtime"
	"os"
	"os/exec"
	"time"
)
func main(){
	var s string = runtime.GOOS
	fmt.Println(s)
	clearTerminal()
	for true {
		fmt.Println(s)
		time.Sleep(5 * time.Millisecond)
		clearTerminal()
		s =  s 
	}
}

func bruteForce(){

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
