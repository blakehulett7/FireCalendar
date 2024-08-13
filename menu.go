package main

import (
	"fmt"
	"os"
)

func DrawMenu() {
	currentIndex := 0
	options := getMenuOptions()
	optionsArray := []string{}
	for option := range options {
		optionsArray = append(optionsArray, option)
	}
	for index, option := range optionsArray {
		if index == currentIndex {
			fmt.Println(cyan + "> " + option + reset)
			continue
		}
		fmt.Println(" ", option)
	}
}

func getMenuOptions() map[string]func() {
	return map[string]func(){
		"Add Event":            AddEvent,
		"View Upcoming Events": ViewEvents,
		"Exit":                 Exit,
	}
}

func AddEvent() {
	fmt.Println("Adding event...")
}

func ViewEvents() {
	fmt.Println("Viewing events...")
}

func Exit() {
	os.Exit(0)
}
