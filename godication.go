//disabled kiss principle
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Process struct {
	ProcessId     int
	ProcessName   string
	ProcessStatus bool
}

func find(processArr *[]Process, processId int) {
	if processId <= len(*processArr) {
		for _, process := range *processArr {
			if process.ProcessId == processId && process.ProcessStatus {
				successText("Found!")
				resultText(process.ProcessName)
				return
			} else if !process.ProcessStatus {
				infoText("Process is currently unavailable")
			}
		}
	} else {
		errorText("Failure.")
	}
}

func exceptionHandler() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Detail", r)
		}
	}()
	panic("Was a problem.")
}

func successText(text string) {
	fmt.Println("Success: " + text)
}
func errorText(text string) {
	fmt.Println("Error: " + text)
}

func infoText(text string) {
	fmt.Println("Info: " + text)
}

func resultText(text string) {
	fmt.Println("Result: " + text)
}

func separateText() {
	fmt.Println("-------------------------------------------------------")
}
func main() {
	//static processes
	processes := []Process{
		{ProcessId: 1, ProcessName: "Hello", ProcessStatus: true},
		{ProcessId: 2, ProcessName: "World", ProcessStatus: true},
		{ProcessId: 3, ProcessName: "Only Education", ProcessStatus: true},
		{ProcessId: 4, ProcessName: "Disabled", ProcessStatus: false},
	}
	//reader that read console
	reader := bufio.NewReader(os.Stdin)
	//loop
	for i := 0; i < 1; {
		fmt.Println("1 : Returns Hello")
		fmt.Println("2 : Returns World")
		fmt.Println("3 : Returns Project Reason")
		fmt.Println("4 : Returns X")
		fmt.Println("Enter your process id:")

		readText, _ := reader.ReadString('\n')
		processId := strings.TrimSpace(readText)
		id, err := strconv.Atoi(processId)
		if err != nil {
			fmt.Println("Please enter a valid number")
		} else {
			find(&processes, id)
		}
		separateText()
	}

}
