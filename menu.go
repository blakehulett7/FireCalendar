package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

type menuOption struct {
	index   int
	name    string
	command func()
}

func MenuInputLoop(currentIndex int) int {
	optionsArray := getMenuOptions()
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	var input []byte = make([]byte, 1)
	for {
		os.Stdin.Read(input)
		//fmt.Println("I got the byte", input, "("+string(input)+")")
		pressedKey := string(input)
		if pressedKey == "j" {
			currentIndex++
			currentIndex = WrapIndex(currentIndex)
			return currentIndex
		}
		if pressedKey == "k" {
			currentIndex--
			currentIndex = WrapIndex(currentIndex)
			return currentIndex
		}
		if pressedKey == "\n" {
			optionsArray[currentIndex].command()
			return currentIndex
		}
	}
}

func WrapIndex(currentIndex int) int {
	arrayLength := len(getMenuOptions())
	if currentIndex < 0 {
		return arrayLength - 1
	}
	if currentIndex > arrayLength-1 {
		return 0
	}
	return currentIndex
}

func DrawMenu(currentIndex int) {
	optionsArray := getMenuOptions()
	if currentIndex < 0 || currentIndex > len(optionsArray)-1 {
		fmt.Println("Error, index out of optionsArray")
		return
	}
	for index, option := range optionsArray {
		if index == currentIndex {
			fmt.Println(" " + cyan + "> " + option.name + reset)
			continue
		}
		fmt.Println("  ", option.name)
	}
}

func getMenuOptions() []menuOption {
	return []menuOption{
		{
			index:   0,
			name:    "Add Event",
			command: AddEvent,
		},
		{
			index:   1,
			name:    "View Upcoming Events",
			command: ViewEvents,
		},
		{
			index:   2,
			name:    "Exit",
			command: Exit,
		},
	}
}

func AddEvent() {
	Run("clear")
	fmt.Println("Adding event...")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
}

func ViewEvents() {
	Run("clear")
	fmt.Println("Viewing events...")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
}

func Exit() {
	Run("tput", "cnorm")
	Run("stty", "-F", "/dev/tty", "echo")
	os.Exit(0)
}
