package main

import (
	"fmt"
	"os"
)

type menuOption struct {
	index   int
	name    string
	command func()
}

func DrawMenu(currentIndex int) {
	optionsArray := getMenuOptions()
	for index, option := range optionsArray {
		if index == currentIndex {
			fmt.Println(cyan + "> " + option.name + reset)
			continue
		}
		fmt.Println(" ", option.name)
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
	fmt.Println("Adding event...")
}

func ViewEvents() {
	fmt.Println("Viewing events...")
}

func Exit() {
	os.Exit(0)
}
