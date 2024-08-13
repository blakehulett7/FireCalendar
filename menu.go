package main

import (
	"fmt"
)

func StartMenu() {
}

func getMenuOptions() map[string]func() {
	return map[string]func(){
		"Add Event":            AddEvent,
		"View Upcoming Events": ViewEvents,
	}
}

func AddEvent() {
	fmt.Println("Adding event...")
}

func ViewEvents() {
	fmt.Println("Viewing events...")
}
