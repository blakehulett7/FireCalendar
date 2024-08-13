package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		clear := exec.Command("clear")
		clear.Stdout = os.Stdout
		clear.Run()
		fmt.Println("Welcome to the FireCalendar, Christ is King!")
		fmt.Println("")
		DrawMenu()
		scanner.Scan()
	}
}

const cyan = "\033[36m"
const reset = "\033[0m"
