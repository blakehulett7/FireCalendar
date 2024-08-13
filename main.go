package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	Run("stty", "-F", "/dev/tty", "-echo")
	Run("tput", "civis")
	currentIndex := 0
	for {
		//Run("clear")
		fmt.Println("Welcome to the FireCalendar, Christ is King!")
		fmt.Println("")
		DrawMenu(currentIndex)
		currentIndex = MenuInputLoop(currentIndex)
	}
}

const cyan = "\033[36m"
const reset = "\033[0m"

func Run(program string, args ...string) {
	command := exec.Command(program, args...)
	command.Stdout = os.Stdout
	command.Run()
}
