package main

import (
	//"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	keypress := make([]byte, 1)
	//scanner := bufio.NewScanner(os.Stdin)
	currentIndex := 0
	for {
		clear := exec.Command("clear")
		clear.Stdout = os.Stdout
		clear.Run()
		fmt.Println("Welcome to the FireCalendar, Christ is King!")
		fmt.Println("")
		DrawMenu(currentIndex)
		os.Stdin.Read(keypress)
		fmt.Println(keypress)
	}
}

const cyan = "\033[36m"
const reset = "\033[0m"
