package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	clearCursor := exec.Command("tput", "civis")
	clearCursor.Stdout = os.Stdout
	clearCursor.Run()
	//keypress := make([]byte, 1)
	scanner := bufio.NewScanner(os.Stdin)
	currentIndex := 0
	for {
		clear := exec.Command("clear")
		clear.Stdout = os.Stdout
		clear.Run()
		fmt.Println("Welcome to the FireCalendar, Christ is King!")
		fmt.Println("")
		DrawMenu(currentIndex)
		//os.Stdin.Read(keypress)
		//fmt.Println(keypress)
		scanner.Scan()
		response := scanner.Text()
		if response == "exit" {
			showCursor := exec.Command("tput", "cnorm")
			showCursor.Stdout = os.Stdout
			showCursor.Run()
			os.Exit(0)
		}
	}
}

const cyan = "\033[36m"
const reset = "\033[0m"
