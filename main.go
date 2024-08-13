package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	Run("stty", "-F", "/dev/tty", "-echo")
	Run("tput", "civis")
	//keypress := make([]byte, 1)
	scanner := bufio.NewScanner(os.Stdin)
	currentIndex := 0
	for {
		Run("clear")
		fmt.Println("Welcome to the FireCalendar, Christ is King!")
		fmt.Println("")
		DrawMenu(currentIndex)
		//os.Stdin.Read(keypress)
		//fmt.Println(keypress)
		scanner.Scan()
		response := scanner.Text()
		if response == "exit" {
			Run("tput", "cnorm")
			Run("stty", "-F", "/dev/tty", "echo")
			os.Exit(0)
		}
	}
}

const cyan = "\033[36m"
const reset = "\033[0m"

func Run(program string, args ...string) {
	command := exec.Command(program, args...)
	command.Stdout = os.Stdout
	command.Run()
}
