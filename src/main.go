package main

import (
	"fmt"
	"os"
)

func main() {

	showMenu()
	command := setCommand()

	switch command {
	case 1:
		fmt.Println("Monitoring")
	case 2:
		fmt.Println("Showing logs")
	case 0:
		fmt.Println("Exiting")
		os.Exit(0)
	default:
		fmt.Println("Invalid option")
	}
}

func showMenu() {
	fmt.Println("Welcome")

	fmt.Println()

	fmt.Println("1- Start Monitoring")
	fmt.Println("2- Show Logs")
	fmt.Println("0- Exit")

	fmt.Println()
}

func setCommand() int {
	var command int
	fmt.Scan(&command)

	return command
}
